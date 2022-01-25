package cmd

import (
	"fmt"
	"github.com/ek/pkg/k8s"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"log"
	"os"
	"reflect"
	"strings"
)

var pCmd = &cobra.Command{
	Use: "p",
	Run: func(cmd *cobra.Command, args []string) {
		pRun()
	},
}

func init() {
	rootCmd.AddCommand(pCmd)
}

type resource struct {
	Name string
}

var actions = []resource {
	{ Name: "logs" },
	{ Name: "describe" },
}

var resourceTypes = []resource {
	{ Name: "sts" },
	{ Name: "deployment" },
	{ Name: "pod" },
	{ Name: "exit"},
}

/**
操作逻辑：
	1、选择资源类型 （sts deployment pod ...）
	2、选择资源（resource name）
	3、选择操作 (logs describe)
*/
func pRun() {
	num := selectUI(resourceTypes)
	resourceTypeName := resourceTypes[num]
	fmt.Printf("resouce type: %v\n", resourceTypeName.Name)

	k8sClient := k8s.NewClient()

	// 获取namespace
	namespaceList := k8sClient.FetchNamespaces()
	var namespaces []resource
	for _, item := range namespaceList.Items {
		namespaces = append(namespaces, resource{
			Name: item.Name,
		})
	}
	namespaces = append(namespaces, resource{
		Name: "exit",
	})
	namespaceIndex := selectUI(namespaces)
	ns := namespaces[namespaceIndex].Name

	// 获取资源
	podList := k8sClient.FetchPods(ns)
	var resources []resource
	for _, item := range podList.Items {
		resources = append(resources, resource {
			Name: item.Name,
		})
	}
	resources = append(resources, resource{
		Name: "exit",
	})

	resourceIndex := selectUI(resources)
	//resourceName := resources[resourceIndex]
	selectPod := podList.Items[resourceIndex]
	//
	// 选择操作
	actionIndex := selectUI(actions)
	actionName := actions[actionIndex].Name

	if actionName == "describe" {
		k8sClient.DescribePod(selectPod)
	}

}

func selectUI(items interface{}) int  {
	values := reflect.ValueOf(items)
	sliceLen := values.Len()
	out := make([]interface{}, sliceLen)
	for i := 0; i < sliceLen; i++ {
		out[i] = values.Index(i).Interface()
	}

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F336 {{ .Name | cyan }}",
		Inactive: "  {{ .Name | cyan }}",
		Selected: "\U0001F336 {{ .Name | red | cyan }}",
		Details: `
--------- ResourceType ----------
{{ "Name:" | faint }}	{{ .Name }}`,
	}

	search := func(input string, index int) bool {
		item := out[index].(resource)

		name := strings.Replace(strings.ToLower(item.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label: "choose resource type",
		Items: items,
		Templates: templates,
		Size: 4,
		Searcher: search,
		HideHelp: true,
	}

	i, name, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	if name == "exit" {
		fmt.Println("Bye bye")
		os.Exit(1)
	}

	return i
}
