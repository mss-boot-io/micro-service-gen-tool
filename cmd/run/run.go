package run

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"

	"github.com/mss-boot-io/micro-service-gen-tool/pkg"
)

var (
	StartCmd = &cobra.Command{
		Use:     "run",
		Short:   "Start generate project",
		Example: "generate-tool run",
		PreRun: func(cmd *cobra.Command, args []string) {
			pkg.Upgrade(true)
			pre()

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	defaultTemplate = "git@github.com:WhiteMatrixTech/matrix-microservice-template.git"
)

func pre() {
	if os.Getenv("MICRO_DEFAULT_TEMPLATE") != "" {
		defaultTemplate = os.Getenv("MICRO_DEFAULT_TEMPLATE")
	}
}

func run() error {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	var err error
	//var repo string
	repo := defaultTemplate
	fmt.Printf("template repo (default:%s): ", pkg.Yellow(defaultTemplate))
	_, _ = fmt.Scanf("%s", &repo)
	fmt.Println("your template repo: ", pkg.Cyan(repo))
	branch := "main"
	fmt.Printf("template repo branch(default:'%s'): ", pkg.Yellow(branch))
	_, _ = fmt.Scanf("%s", &branch)
	home, err := homedir.Dir()
	privateID := "id_rsa"
	fmt.Printf("private key id(default:%s): ", pkg.Yellow(privateID))
	_, _ = fmt.Scanf("%s", &privateID)
	privateKeyFile := filepath.Join(home, ".ssh", privateID)
	if err != nil {
		log.Fatalln(err)
	}
	templateWorkspace := "/tmp/template-workspace"
	fmt.Printf("template workspace(default:%s): ", pkg.Yellow(templateWorkspace))
	_, _ = fmt.Scanf("%s", &templateWorkspace)
	templateWorkspace = filepath.Join(templateWorkspace, uuid.New().String())
	var password string
	fmt.Printf("private pem password(default:%s): ", pkg.Yellow("''"))
	_, _ = fmt.Scanf("%s", &password)
	fmt.Printf("git clone start: %s \n", time.Now().String())
	fmt.Println(privateKeyFile)
	err = pkg.GitCloneSSH(repo, templateWorkspace, branch, privateKeyFile, password)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("git clone end: %s \n", time.Now().String())
	_ = os.RemoveAll(filepath.Join(templateWorkspace, ".git"))
	defer os.RemoveAll(templateWorkspace)
	sub, err := pkg.GetSubPath(templateWorkspace)
	if err != nil {
		log.Fatalln(err)
	}
	if len(sub) == 0 {
		log.Fatalln("not found template")
	}
	fmt.Println(pkg.Yellow("************* please select sub path ***************"))
	for i := range sub {
		fmt.Println(pkg.Yellow("* "), pkg.Yellow(sub[i]))
	}
	fmt.Println(pkg.Yellow("****************************************************"))
SUBPATH:
	subPath := sub[0]
	fmt.Printf("select template sub path(default:%s): ", pkg.Cyan(subPath))
	_, _ = fmt.Scanf("%s", &subPath)
	ok := false
	for i := range sub {
		if sub[i] == subPath {
			ok = true
		}
	}
	if !ok {
		fmt.Println(pkg.Red("please select exist sub path."))
		goto SUBPATH
	}
	projectName := "default"
	fmt.Printf("project name(default:%s)", pkg.Yellow(projectName))
	_, _ = fmt.Scanf("%s", &projectName)
	keys, err := pkg.GetParseFromTemplate(templateWorkspace, subPath)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(pkg.Magenta("please config your param's value"))

	for key := range keys {
		var value string
	BACK:
		fmt.Printf("%s: ", pkg.Yellow(key))
		_, _ = fmt.Scanf("%s", &value)
		if value == "" {
			goto BACK
		}
		keys[key] = value
	}

	err = pkg.Generate(&pkg.TemplateConfig{
		Service:       subPath,
		TemplateLocal: templateWorkspace,
		CreateRepo:    false,
		Destination:   filepath.Join(".", projectName),
		Github:        nil,
		Params:        keys,
		Ignore:        nil,
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pkg.Green("template generate project success...."))
	return nil
}
