package vfs

import (
	"fmt"
	"path"
	"sort"
)

// List возвращает имя файла или отсортированное содержимое каталога.
// Имена каталогов заканчиваются символом /.
func (fs *FileSystem) List(name string) ([]string, error) {
	target, err := fs.Resolve(name)
	if err != nil {
		return nil, err
	}

	entry, exists := fs.Entries[target]
	if !exists {
		return nil, fmt.Errorf("путь не существует: %s", target)
	}
	if !entry.IsDir {
		return []string{path.Base(target)}, nil
	}

	names := make([]string, 0)
	for entryPath, child := range fs.Entries {
		if entryPath == target || path.Dir(entryPath) != target {
			continue
		}

		childName := path.Base(entryPath)
		if child.IsDir {
			childName += "/"
		}
		names = append(names, childName)
	}

	sort.Strings(names)
	return names, nil
}
