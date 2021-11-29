//This program checks a directory input by the user, goes through all the subdirectories
//and checks all files for duplicates. Duplicate files are files that share name and size.
//
//The scanDir method
//
//Asks the user for a path to a directory.
//
//  func (df *dirFiles) scanDir() error {
//  	fmt.Println("Please, Enter directory to scan for duplicate files.")
//
//  	_, err := fmt.Scan(&df.dir)
//  	if err != nil {
//  		return err
//  	}
//  	err = os.Chdir(df.dir)
//  	if err != nil {
//  		return err
//  	}
//
//  	log.Println(`Your directory is`, df.dir)
//  	return nil
//  }
//
//The walkDir method
//
//Creates a slice of strings with paths to all files in the chosen directory.
//
//  func (df *dirFiles) walkDir() error {
//  	err := filepath.Walk(df.dir, func(path string, info os.FileInfo, err error) error {
//  		if !info.IsDir() {
//  			df.files = append(df.files, path)
//  		}
//  		return nil
//  	})
//  	return err
//  }
//The findDuplicates method
//
//Checks all the files names and sizes for duplicates.
//If a duplicate files is found, it is removed from the slice of files
//and added to a slice of strings with paths to duplicate files. If no duplicates are found, program exits.
//
//  func (df *dirFiles) findDuplicates() error {
//  	for i := 0; i < len(df.files); i++ {
//  		iInfo, err := os.Stat(df.files[i])
//  		if err != nil {
//  			log.Printf("Couldn't get %s stats", df.files[i])
//  			return err
//  		}
//  		for j := 1; j < len(df.files); j++ {
//  			jInfo, err := os.Stat(df.files[j])
//  			if err != nil {
//  				log.Printf("Couldn't get %s stats", df.files[j])
//  				return err
//  			}
//  			if iInfo.Name() == jInfo.Name() && iInfo.Size() == jInfo.Size() && df.files[i] != df.files[j] {
//  				df.duplicates = append(df.duplicates, df.files[j])
//  				df.files[j] = df.files[len(df.files)-1]
//  				df.files[len(df.files)-1] = ""
//  				df.files = df.files[:len(df.files)-1]
//  				j = 1
//  			}
//  		}
//  	}
//  	if len(df.duplicates) == 0 {
//  		log.Println("No duplicate files found. Exiting..")
//  		os.Exit(1)
//  	} else {
//  		for _, f := range df.files {
//  			log.Printf(`File "%s" has duplicates`, f)
//  		}
//  		log.Printf("Duplicate files: %v", df.duplicates)
//  	}
//  	return nil
//  }
//The deleteDuplicates method
//
//Concurently goes through all duplicate files and deletes them at user request.
//
//  func (df *dirFiles) deleteDuplicates(ctx context.Context) error {
//  	var answer string
//
//  	if *delFlag {
//  		answer = "y"
//  	} else {
//  		fmt.Println("Do you wish to delete all duplicates? y/n")
//  		_, err := fmt.Scan(&answer)
//  		for err != nil {
//  			return err
//  		}
//  	}
//  	switch answer {
//  	case "y":
//  		for _, f := range df.duplicates {
//  			df.wg.Add(1)
//  			go func(ff string) {
//  				df.mu.Lock()
//  				defer df.mu.Unlock()
//  				err := os.Remove(ff)
//  				if err != nil {
//  					log.Println(checkError("Couldn't delete file:"), ff)
//  				}
//  				df.wg.Done()
//  			}(f)
//  		}
//  		df.wg.Wait()
//
//  		log.Println("All duplicate files successfully deleted.")
//  	case "n":
//  		log.Println("All duplicate files remain.")
//  	default:
//  		log.Println(`Please, answer "y" or "n".`)
//  		err := checkError("Wrong input by user.")
//  		return err
//  	}
//  	return nil
//  }
//The copyOriginals method
//
//Concurrently copies all original files to a new directory called "originals" at user request.
//
//  func (df *dirFiles) copyOriginals(ctx context.Context) error {
//  	var answer string
//
//  	if *copyFlag {
//  		answer = "y"
//  	} else {
//  		fmt.Println(`Do you wish to move all original files to a new directory "*your_dir*/originals"? y/n`)
//  		_, err := fmt.Scan(&answer)
//  		for err != nil {
//  			return err
//  		}
//  	}
//  	switch answer {
//  	case "y":
//  		err := os.Mkdir(df.dir+"/originals", os.ModePerm)
//  		if err != nil {
//  			log.Println(`Couldn't create directory "originals"`)
//  			return err
//  		}
//  		for _, f := range df.files {
//  			df.wg.Add(1)
//  			go func(ff string) {
//  				df.mu.Lock()
//  				defer df.mu.Unlock()
//  				fInfo, err := os.Stat(ff)
//  				if err != nil {
//  					log.Println(checkError("Couldn't get stats of file:"), ff)
//  				}
//  				err = os.Rename(ff, df.dir+"/originals/"+fInfo.Name())
//  				if err != nil {
//  					log.Println(checkError("Couldn't move original file:"), ff)
//  				}
//  				df.wg.Done()
//  			}(f)
//  		}
//  		df.wg.Wait()
//
//  		log.Println("Files successfully moved to the new location.")
//  	case "n":
//  		log.Println("Program exited at user request.")
//  	default:
//  		log.Println(`Please, answer "y" or "n".`)
//  		err := checkError("Wrong input by user.")
//  		return err
//  	}
//  	return nil
//  }
package doc
