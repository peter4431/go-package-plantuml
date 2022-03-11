package main

import (
	"fmt"
	"git.oschina.net/jscode/go-package-plantuml/codeanalysis"
	"github.com/jessevdk/go-flags"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
)

func main() {

	log.SetLevel(log.InfoLevel)

	var opts struct {
		CodeDir    string   `long:"codedir" description:"要扫描的代码目录" required:"true"`
		OutputFile string   `long:"outputfile" description:"解析结果保存到该文件中" required:"true"`
		IgnoreDirs []string `long:"ignoredir" description:"需要排除的目录,不需要扫描和解析"`
	}

	if len(os.Args) == 1 {
		fmt.Println("使用例子\n" +
			os.Args[0] + " --codedir /appdev/gopath/src/github.com/contiv/netplugin --gopath /appdev/gopath --outputfile  /tmp/result")
		os.Exit(1)
	}

	_, err := flags.ParseArgs(&opts, os.Args)

	if err != nil {
		os.Exit(1)
	}

	if opts.CodeDir == "" {
		panic("代码目录不能为空")
	}

	if !path.IsAbs(opts.CodeDir) {
		wd, _ := os.Getwd()
		opts.CodeDir = path.Join(wd, opts.CodeDir)
	}

	var newDir string
	var newIgnoreDirs []string
	for _, dir := range opts.IgnoreDirs {
		if !path.IsAbs(dir) {
			newDir = path.Join(opts.CodeDir, dir)
		} else {
			newDir = dir
		}
		newIgnoreDirs = append(newIgnoreDirs, newDir)
	}

	config := codeanalysis.Config{
		CodeDir:    opts.CodeDir,
		GopathDir:  opts.CodeDir,
		VendorDir:  path.Join(opts.CodeDir, "vendor"),
		IgnoreDirs: newIgnoreDirs,
	}

	result := codeanalysis.AnalysisCode(config)

	result.OutputToFile(opts.OutputFile)
}
