package build

import (
	"bytes"
	"cmp"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	"github.com/ajr-cabbage/sssg/src/internal/config"
	cp "github.com/nmrshll/go-cp"
	"github.com/yuin/goldmark"
)

type directoryNode struct {
	name     string
	path     string
	isDir    bool
	children []*directoryNode
	mu       sync.Mutex
}

type document struct {
	title string
	path  string
	html  string
}

type documentStore struct {
	docs []document
	mu   sync.Mutex
}

func Build(c config.Config) error {
	// verify config
	ok := config.VerifyConfig(c)
	if !ok {
		return errors.New("Invalid or missing configuration\n'sssg init' to initialize a project in the current directory\n")
	}
	// scan directory
	docStore, dirTree, err := scanDirectory(c.SourcePath, c.TargetPath)
	if err != nil {
		return fmt.Errorf("Fatal error scanning source directory: %w\n", err)
	}
	// write documents
	err = writeDocuments(c, docStore.docs, dirTree)
	if err != nil {
		return fmt.Errorf("Fatal error writing documents: %w\n", err)
	}
	// copy assets
	err = copyAssets(c.AssetPath, filepath.Join(c.ThemePath, c.Theme), filepath.Join(c.TargetPath, "assets"))
	if err != nil {
		return fmt.Errorf("Fatal error copying assets: %w\n", err)
	}
	return nil
}

func scanDirectory(sourceDir, targetDir string) (*documentStore, *directoryNode, error) {
	_, err := os.Stat(targetDir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(targetDir, 0744)
			if err != nil {
				return &documentStore{}, &directoryNode{}, err
			}
		} else {
			return &documentStore{}, &directoryNode{}, err
		}
	}
	sourceDirContents, err := os.ReadDir(sourceDir)
	if err != nil {
		return &documentStore{}, &directoryNode{}, err
	}
	docs := documentStore{
		docs: []document{},
	}
	dirTree := directoryNode{
		name:  "root",
		path:  targetDir,
		isDir: true,
	}
	var wg sync.WaitGroup
	for _, item := range sourceDirContents {
		wg.Go(func() {
			itemSourcePath := filepath.Join(sourceDir, item.Name())
			itemTargetPath := filepath.Join(targetDir, item.Name())
			if item.IsDir() {
				subDocs, subTree, err := scanDirectory(itemSourcePath, itemTargetPath)
				if err != nil {
					fmt.Println(err)
					return
				}
				subTree.isDir = true
				subTree.name = item.Name()
				subTree.path = itemTargetPath
				dirTree.mu.Lock()
				dirTree.children = append(dirTree.children, subTree)
				dirTree.mu.Unlock()
				subDocs.mu.Lock()
				for _, doc := range subDocs.docs {
					docs.mu.Lock()
					docs.docs = append(docs.docs, doc)
					docs.mu.Unlock()
				}
				subDocs.mu.Unlock()
			} else {
				ext := filepath.Ext(itemSourcePath)
				if strings.ToLower(ext) != ".md" {
					fmt.Printf("Invalid file ext {%s} for: %s", ext, itemSourcePath)
					return
				}
				doc, err := processDocument(itemSourcePath, itemTargetPath[:len(itemTargetPath)-len(ext)])
				if err != nil {
					fmt.Println(err)
					return
				}
				docs.mu.Lock()
				docs.docs = append(docs.docs, doc)
				docs.mu.Unlock()
				if item.Name() != "index.md" {
					fileNode := directoryNode{
						name:  doc.title,
						path:  itemTargetPath[:len(itemTargetPath)-len(ext)] + ".html",
						isDir: false,
					}
					dirTree.mu.Lock()
					dirTree.children = append(dirTree.children, &fileNode)
					dirTree.mu.Unlock()
				}
			}
		})
	}
	wg.Wait()
	slices.SortFunc(dirTree.children, func(a, b *directoryNode) int {
		return cmp.Compare(a.name, b.name)
	})
	return &docs, &dirTree, nil
}

func processDocument(documentPath, targetPath string) (document, error) {
	md, err := os.ReadFile(documentPath)
	if err != nil {
		return document{}, fmt.Errorf("Error reading document '%s': %w\n", documentPath, err)
	}
	title := getTitle(string(md))
	html, err := mdToHTML(md)
	if err != nil {
		return document{}, fmt.Errorf("Error converting md to html: %w", err)
	}
	doc := document{
		title: title,
		path:  targetPath + ".html",
		html:  html,
	}
	return doc, nil
}

func mdToHTML(md []byte) (string, error) {
	var buf bytes.Buffer
	err := goldmark.Convert(md, &buf)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func getTitle(md string) string {
	for line := range strings.SplitSeq(md, "\n") {
		if len(line) >= 3 && line[:2] == "# " {
			return line[2:]
		}
		if len(line) >= 4 && line[:3] == "## " {
			return line[3:]
		}
	}
	return "Untitled"
}

func writeDocuments(c config.Config, documents []document, tree *directoryNode) error {
	template, err := os.ReadFile(filepath.Join(c.TemplatePath, "shitty-default.html"))
	if err != nil {
		return fmt.Errorf("Error reading HTML template: %w", err)
	}
	treeString := printTree(tree, c)
	if treeString == "" {
		treeString += `<li><a href="/">HOME</a></li>`
	}
	var wg sync.WaitGroup
	for _, doc := range documents {
		wg.Go(func() {
			htmlDoc := strings.Replace(string(template), "{{ Site Name }}", c.SiteName, 1)
			htmlDoc = strings.Replace(htmlDoc, "{{ Title }}", doc.title, 1)
			htmlDoc = strings.Replace(htmlDoc, "{{ Sidebar }}", treeString, 1)
			htmlDoc = strings.Replace(htmlDoc, "{{ Content }}", doc.html, 1)
			htmlDoc = strings.Replace(htmlDoc, "href=\"/", fmt.Sprintf("href=\"%s", c.BasePath), -1)
			htmlDoc = strings.Replace(htmlDoc, "src=\"/", fmt.Sprintf("src=\"%s", c.BasePath), -1)
			err = os.WriteFile(doc.path, []byte(htmlDoc), 0644)
			if err != nil {
				fmt.Printf("Error writing '%s': %s", doc.path, err.Error())
			}
		})
	}
	wg.Wait()
	return nil
}

func printTree(node *directoryNode, c config.Config) string {
	var treeString string
	if !node.isDir {
		trimmedPath, ok := strings.CutPrefix(node.path, c.TargetPath[2:])
		if ok {
			treeString += fmt.Sprintf(`<li><a href="/%s">%s</a></li>`, trimmedPath, node.name)
		} else {
			treeString += fmt.Sprintf(`<li><a href="/%s">%s</a></li>`, node.path, node.name)
		}
		return treeString
	}
	if node.children == nil && node.name != "root" {
		treeString += fmt.Sprintf(`<h2 class="sidebar__heading">%s</h2>`, node.name)
		return treeString
	}
	if node.name != "root" {
		treeString += fmt.Sprintf(`<h2 class="sidebar__heading">%s</h2>`, node.name)
	}
	node.mu.Lock()
	for _, child := range node.children {
		treeString += printTree(child, c)
	}
	node.mu.Unlock()
	return treeString
}

func copyAssets(assetSource, cssSource, assetTarget string) error {
	_, err := os.Stat(assetTarget)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = os.MkdirAll(assetTarget, 0744)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	sourceDirContents, err := os.ReadDir(assetSource)
	if err != nil {
		return err
	}
	for _, item := range sourceDirContents {
		itemSourcePath := filepath.Join(assetSource, item.Name())
		itemTargetPath := filepath.Join(assetTarget, item.Name())
		if item.IsDir() {
			copyAssets(itemSourcePath, itemTargetPath, cssSource)
		} else {
			err := cp.CopyFile(itemSourcePath, itemTargetPath)
			if err != nil {
				fmt.Printf("Copy file error: %s\n", err.Error())
			}
		}
	}
	// Copy theme css file to target assset directory.
	// Print, but do not return errors. Not critical if theme is not copied.
	_, err = os.Stat(cssSource)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("Theme missing: %s\n", cssSource)
			return nil
		}
		fmt.Printf("Error reading theme: %s\n", err.Error())
		return nil
	}
	err = cp.CopyFile(cssSource, filepath.Join(assetTarget, "index.css"))
	if err != nil {
		fmt.Printf("Error copying theme: %s\n", err.Error())
	}

	return nil
}
