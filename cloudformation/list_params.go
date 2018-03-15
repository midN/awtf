package cloudformation

import (
	"errors"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"io"
	"sort"
	"strings"
)

type LpParams struct {
	StackName string
	Sess      *session.Session
	TableType string
}

type tableData [][]string

func (c tableData) Len() int      { return len(c) }
func (c tableData) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

// Compare strings by first key of the sub-slice
func (c tableData) Less(i, j int) bool { return strings.Compare(c[i][0], c[j][0]) == -1 }

func ListParams(lpParams LpParams, writer io.Writer) error {
	if lpParams.StackName == "" {
		return errors.New("Stack name is missing!")
	}

	cfSession := cloudformation.New(lpParams.Sess)
	params := cloudformation.DescribeStacksInput{StackName: aws.String(lpParams.StackName)}
	stackInfo, err := cfSession.DescribeStacks(&params)

	if err != nil {
		return err
	}

	stackParams := stackInfo.Stacks[0].Parameters
	data := sortedStackParams(stackParams)
	printTable(data, writer, lpParams.TableType)

	return nil
}

func printTable(data tableData, writer io.Writer, tableType string) {
	table := tablewriter.NewWriter(writer)

	switch tableType {
	case "markdown":
		setMarkdownTableAttributes(table)
	default:
		setDefaultTableAttributes(table)
	}

	setSharedTableAttributes(table)

	table.AppendBulk(data)
	table.Render()
}

func setMarkdownTableAttributes(table *tablewriter.Table) {
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
}

func setDefaultTableAttributes(table *tablewriter.Table) {
	table.SetRowLine(true)
}

func setSharedTableAttributes(table *tablewriter.Table) {
	titleColor := color.New(color.FgCyan).SprintFunc()
	table.SetAutoFormatHeaders(false)
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	table.SetHeader([]string{titleColor("Parameter Name"), titleColor("Parameter Value")})
}

func sortedStackParams(stackParams []*cloudformation.Parameter) tableData {
	keyColor := color.New(color.FgGreen, color.Bold).SprintfFunc()
	valueColor := color.New(color.FgBlue).SprintfFunc()

	data := tableData{}
	for _, v := range stackParams {
		data = append(data, []string{keyColor(*v.ParameterKey), valueColor(*v.ParameterValue)})
	}
	sort.Sort(data)

	return data
}
