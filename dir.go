package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func showDirWindow(inputField *tview.InputField) {
	windowName := "showDirWindow"
	rootDir := "/"

	node := tview.NewTreeNode(rootDir).SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(node).
		SetCurrentNode(node)

	tree.SetBorder(true).SetTitle("Directory")

	// A helper function which adds the files and directories of the given path
	// to the given target node.
	add := func(target *tview.TreeNode, path string) {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			node := tview.NewTreeNode(file.Name()).
				SetReference(filepath.Join(path, file.Name())).
				SetSelectable(file.IsDir())
			if file.IsDir() {
				node.SetColor(tcell.ColorGreen)
			}
			target.AddChild(node)
		}
	}

	add(node, rootDir)

	// path route
	var route []string
	// recursive to root
	dir := inputField.GetText()

	// get file info
	fi, _ := os.Stat(dir)

	for dir != rootDir {
		if dir == "." {
			wd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			dir = wd
		}

		dir = filepath.Dir(dir)
		route = append(route, dir)
	}

	// create a path to inputfield
	prevNode := node
	for i := len(route) - 1; i >= 0; i-- {
		children := prevNode.GetChildren()
		for _, children := range children {
			dir := children.GetReference().(string)
			if dir == route[i] {
				tree.SetCurrentNode(children)
				files, err := ioutil.ReadDir(dir)
				if err != nil {
					panic(err)
				}

				for _, file := range files {
					node := tview.NewTreeNode(file.Name()).
						SetReference(filepath.Join(route[i], file.Name())).
						SetSelectable(true)
					if file.IsDir() {
						node.SetColor(tcell.ColorGreen)
					} else {
						// select the given file
						if os.SameFile(fi, file) {
							tree.SetCurrentNode(node)
						}
					}
					children.AddChild(node)
				}

				children.Expand()
				prevNode = children
				break
			}
		}
	}

	// Change input text based on tree selection
	tree.SetChangedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		path := reference.(string)
		// set this path to input text
		inputField.SetText(path)
	})

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		path := reference.(string)

		// check if the selected path is dir
		fi, err := os.Stat(inputField.GetText())
		if err != nil {
			panic(err)
		}

		if fi.IsDir() {
			if len(children) == 0 {
				// Load and show files in this directory.
				add(node, path)
			} else {
				// Collapse if visible, expand if collapsed.
				node.SetExpanded(!node.IsExpanded())
			}
		} else {
			// file selected, close window
			root.RemovePage(windowName)
		}
	})

	root.AddPage(windowName, popup(80, 30, tree), true, true)
}
