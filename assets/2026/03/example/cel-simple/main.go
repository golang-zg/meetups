package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/common/decls"
	"github.com/google/cel-go/ext"
	"gopkg.in/yaml.v3"
)

type Applicant struct {
	Email            string `json:"email"`
	PercentagePassed int    `json:"percentagePassed"`
}

type ValidationRule struct {
	Name       string `yaml:"name"`
	Expression string `yaml:"expression"`
	Message    string `yaml:"message"`
}

type Schema struct {
	CELValidations []ValidationRule `yaml:"celValidations"`
}

type Root struct {
	Applicant Schema `yaml:"applicant"`
}

func loadSchema(path string) (Schema, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Schema{}, err
	}

	var root Root
	if err := yaml.Unmarshal(data, &root); err != nil {
		return Schema{}, err
	}

	return root.Applicant, nil
}

func ValidateApplicant(a Applicant, schema Schema) error {
	env, err := cel.NewEnv(
		ext.NativeTypes(reflect.TypeOf(time.Time{})),
		ext.Strings(),
		ext.Math(),
		cel.VariableDecls(
			decls.NewVariable("email", cel.StringType),
			decls.NewVariable("percentagePassed", cel.IntType),
		),
	)
	if err != nil {
		return err
	}

	for _, rule := range schema.CELValidations {
		ast, issues := env.Compile(rule.Expression)
		if issues != nil && issues.Err() != nil {
			return issues.Err()
		}

		prg, err := env.Program(ast)
		if err != nil {
			return err
		}

		out, _, err := prg.Eval(map[string]any{
			"email":            a.Email,
			"percentagePassed": a.PercentagePassed,
		})
		if err != nil {
			return err
		}

		ok := out.Value().(bool)
		if !ok {
			return errors.New(rule.Message)
		}
	}

	return nil
}

func main() {
	schema, err := loadSchema("validation.yaml")
	if err != nil {
		panic(err)
	}

	applicant := Applicant{
		Email:            "test@example.com",
		PercentagePassed: 50,
	}

	if err := ValidateApplicant(applicant, schema); err != nil {
		fmt.Println("Validation failed:", err)
	} else {
		fmt.Println("Validation passed!")
	}
}
