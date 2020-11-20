package cmd

import (
    "github.com/spf13/cobra"
    "strings"
    "github.com/nezhahh/dorakit/internal/fio"
)

var fioFile string
var fioDir string
var outputXlsx string

var desc = strings.Join([]string{
    "分析fio的测试结果文件，汇总数据到xlsx表格"
}, "\n")

var fioCmd = &cobra.Command{
    Use:    "fio",
    Short:  "fio结果分析汇总",
    Long:   desc,
    Run:    func(cmd *cobra.Command, args []string) {
        var fioArgs = FioCmdArgs{
            in:     fioDir
            out:    outputXlsx
            print:  false
        }
    },
}

// dorakit fio
//        -f file
//        -d dir
//        -w xlsx
//        -p print
func init() {
    // TODO: -f和-d命令是互斥的，怎么解决？
    fioCmd.Flags().StringVarP(&fioFile, "file", "f", "", "请输入fio结果文件")
    fioCmd.Flags().StringVarP(&fioDir, "directory", "d", "", "请输入fio结果文件夹")
    fioCmd.Flags().StringVarP(&outputXlsx, "xlsx", "w", "", "请输入结果汇总文件，不设置时，默认在当前命令下以时间戳为名")
}