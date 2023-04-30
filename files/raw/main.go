package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"time"
)

func read_file(file_name string) string {
	file_content, err := os.ReadFile(file_name)
	if err != nil {
		panic(err)
	}
	return string(file_content)
}

func append_file(file_name string, file_content string) {
	f, err := os.OpenFile(file_name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if _, err := f.WriteString("\n" + file_content); err != nil {
		panic(err)
	}
}

func create_file(file_name string, file_content []byte) {
	err := os.WriteFile(file_name, file_content, 0644)
	if err != nil {
		panic(err)
	}
}

func create_temp_file(folder, file_name string) string {
	file, err := os.CreateTemp(folder, file_name+"-")
	if err != nil {
		panic(err)
	}
	return file.Name()
}

func delete_file(file_name string) {
	err := os.Remove(file_name)
	if err != nil {
		panic(err)
	}
}

func move_file(old_path, new_path string) {
	err := os.Rename(old_path, new_path)
	if err != nil {
		panic(err)
	}
}

func rename_file(old_path, new_path string) {
	err := os.Rename(old_path, new_path)
	if err != nil {
		panic(err)
	}
}

func copy_file(src_file, dest_file string) (int64, error) {
	src_file_stat, err := os.Stat(src_file)
	if err != nil {
		panic(err)
	}
	if !src_file_stat.Mode().IsRegular() {
		fmt.Printf("%s is not a regular file", src_file)
	}
	open_src_file, err := os.Open(src_file)
	if err != nil {
		panic(err)
	}
	defer open_src_file.Close()

	dest_file_creation, err := os.Create(dest_file)
	if err != nil {
		panic(err)
	}
	defer dest_file_creation.Close()

	nbytes, err := io.Copy(dest_file_creation, open_src_file)
	if err != nil {
		panic(err)
	}

	return nbytes, err
}

func truncate_file(file_name string) {
	err := os.Truncate(file_name, 0)
	if err != nil {
		panic(err)
	}
}

func check_if_file_exists(file_name string) bool {
	_, err := os.Stat(file_name)
	var file_status bool
	if err == nil {
		file_status = true
	} else {
		file_status = false
	}
	return file_status
}

type FileStat struct {
	Name    string
	IsDir   bool
	Mode    fs.FileMode
	ModTime time.Time
	Size    int64
	Sys     any
}

func check_file_stats(file_name string) FileStat {
	stats, err := os.Stat(file_name)
	if err != nil {
		panic(err)
	}

	file_stats := FileStat{
		Name:    stats.Name(),
		IsDir:   stats.IsDir(),
		Mode:    stats.Mode(),
		ModTime: stats.ModTime(),
		Size:    stats.Size(),
		Sys:     stats.Sys(),
	}

	return file_stats
}

func main() {
	// ---------- READ:FILE[RAW] ----------
	read_file_name := "data/file_to_be_readed.txt"
	fmt.Println(read_file(read_file_name))

	// ---------- APPEND:FILE[RAW] ----------
	append_file_name := "data/file_to_be_appended.txt"
	append_file_content := "some text appended"
	append_file(append_file_name, append_file_content)
	fmt.Println(read_file(append_file_name))

	// ---------- CREATE:FILE[RAW] ----------
	create_file_name := "data/file_to_be_created.txt"
	create_file_content := []byte("----- FILE:CREATE[RAW] -----\nhello, i'm a file to be created!")
	create_file(create_file_name, create_file_content)
	fmt.Println(read_file(create_file_name))

	// ---------- CREATE_TEMP:FILE[RAW] ----------
	create_temp_file_name := "temp_file_to_be_created"
	fmt.Printf("----- CREATE_TEMP:FILE[RAW] -----\n%s", create_temp_file("data", create_temp_file_name))

	// ---------- DELETE:FILE[RAW] ----------
	delete_file_name := create_temp_file("data", "some_temp_file")
	delete_file(delete_file_name)
	fmt.Printf("\n----- DELETE:FILE[RAW] -----\n%s", delete_file_name)

	// ---------- MOVE:FILE[RAW] ----------
	move_file_name := create_temp_file("data", "file_to_be_moved_before.txt")
	move_file(move_file_name, "data/file_to_be_moved_after.txt")
	fmt.Printf("\n----- MOVE:FILE[RAW] -----\n%s", move_file_name)

	// ---------- RENAME:FILE[RAW] ----------
	rename_file_name := create_temp_file("data", "file_to_be_renamed_before.txt")
	rename_file(rename_file_name, "data/file_to_be_renamed_after.txt")
	fmt.Printf("\n----- RENAME:FILE[RAW] -----\n%s", rename_file_name)

	// ---------- COPY:FILE[RAW] ----------
	copy_temp_file_src := create_temp_file("data", "src_file.txt")
	copy_temp_file_dest := create_temp_file("data", "dest_file.txt")
	fmt.Printf("\n----- COPY:FILE[RAW] -----\n")
	_, err := copy_file(copy_temp_file_src, copy_temp_file_dest)
	if err != nil {
		panic(err)
	}
	fmt.Println(copy_temp_file_dest)

	// ---------- TRUNCATE:FILE[RAW] ----------
	truncate_file_name := "data/file_to_be_truncated.txt"
	fmt.Printf("----- TRUNCATE:FILE[RAW] -----\n")
	create_file(truncate_file_name, []byte("some contante to be truncated"))

	fmt.Printf("BEFORE truncate\n%s", read_file(truncate_file_name))
	truncate_file(truncate_file_name)
	fmt.Printf("\nAFTER truncate\n%s", read_file(truncate_file_name))

	// ---------- CHECK_IF_EXISTS:FILE[RAW] ----------
	check_file_exists := create_temp_file("data", "some_file.txt")
	check_exists := check_if_file_exists(check_file_exists)
	fmt.Printf("----- CHECK_IF_EXISTS:FILE[RAW] -----\nfile_status: %t", check_exists)

	// ---------- CHECK_STATS:FILE[RAW] ----------
	check_file_exits_filename := create_temp_file("data", "some_file.txt")
	check_stats := check_file_stats(check_file_exits_filename)
	formated_stats, _ := json.MarshalIndent(&check_stats, "", "  ")
	fmt.Printf("\n----- CHECK_STATS:FILE[RAW] -----\nfile_status: %s", formated_stats)
}
