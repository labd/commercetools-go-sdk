package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	yaml "gopkg.in/mikefarah/yaml.v2"
)

type RamlType struct {
	Name                string
	TypeName            string
	TypeObject          *RamlType
	CodeName            string
	InterfaceName       string
	Discriminator       string
	DiscriminatorObject *RamlType
	DiscriminatorValue  string
	Attributes          []RamlTypeAttribute
	Children            []*RamlType
	Package             string
	resolved            bool
	asMap               bool
	EnumValues          []string
}

type RamlTypeAttribute struct {
	Name        string
	TypeName    string
	TypeObject  *RamlType
	Format      string
	CodeName    string
	Optional    bool
	Many        bool
	Items       string
	ItemsObject *RamlType
}

func (r *RamlType) resolve(types map[string]*RamlType) {
	if r.resolved {
		return
	}
	r.resolved = true

	if value, isOk := types[r.TypeName]; isOk {
		r.TypeObject = value
		value.resolve(types)
		value.Children = append(value.Children, r)
	}

	for i := range r.Attributes {
		r.Attributes[i].resolve(types)
	}
}

func (r *RamlType) isInterface() bool {
	// TODO; check if 1 attribute and that is the discriminator
	if len(r.Attributes) == 1 {
		for _, attr := range r.Children {
			if attr.Discriminator != "" {
				return true
			}
		}
	}
	return false
}

func (r *RamlType) discriminatorFunctionName() string {
	return "mapDiscriminator" + r.InterfaceName
}

func (r *RamlType) getAllAttributes() (result []RamlTypeAttribute) {
	parent := r
	temp := []RamlTypeAttribute{}
	seen := map[string]bool{}
	for {
		for _, attribute := range parent.Attributes {
			if _, found := seen[attribute.CodeName]; found == false {
				temp = append(temp, attribute)
				seen[attribute.CodeName] = true
			}
		}
		parent = parent.TypeObject
		if parent == nil {
			break
		}
	}

	for i := len(temp) - 1; i >= 0; i-- {
		result = append(result, temp[i])
	}

	return result
}

func (r *RamlType) getAttribute(name string) (*RamlTypeAttribute, error) {
	parent := r
	for {
		for _, attribute := range parent.Attributes {
			if attribute.Name == name {
				return &attribute, nil
			}
		}

		parent = parent.TypeObject
		if parent == nil {
			break
		}
	}
	return nil, fmt.Errorf("No attribute %s found on %s", name, r.Name)
}

func (r *RamlType) getDiscriminatorAttribute() *RamlTypeAttribute {
	parent := r

	for {
		if parent.Discriminator != "" {
			attribute, err := r.getAttribute(parent.Discriminator)
			if err != nil {
				panic(err)
			}
			return attribute
		}
		parent = parent.TypeObject
		if parent == nil {
			break
		}
	}

	return nil
}

func (r *RamlType) getDiscriminatedAttributes() (result []*RamlTypeAttribute) {

	mapFields := []string{}
	for _, attr := range r.Attributes {
		if attr.TypeObject != nil && attr.TypeObject.Discriminator != "" {
			mapFields = append(mapFields, attr.Name)
		}

		if attr.ItemsObject != nil && attr.ItemsObject.Discriminator != "" {
			mapFields = append(mapFields, attr.Name)
		}
	}

	for _, name := range mapFields {
		attr, err := r.getAttribute(name)
		if err != nil {
			panic(err)
		}
		result = append(result, attr)
	}
	return result
}

func (r *RamlTypeAttribute) resolve(types map[string]*RamlType) {

	if strings.ContainsAny(r.TypeName, "|") {
		items := []string{}
		for _, item := range strings.Split(r.TypeName, "|") {
			items = append(items, strings.TrimSpace(item))
		}

		if len(items) > 0 {
			fmt.Println(items)
			obj := types[items[0]]
			if obj != nil && obj.TypeObject != nil {
				r.TypeObject = obj.TypeObject
			} else {
				r.TypeObject = obj
			}

		}

	} else if value, isOk := types[r.TypeName]; isOk {
		r.TypeObject = value
	}
	if value, isOk := types[r.Items]; isOk {
		r.ItemsObject = value
	}
}

func getPropertyValue(input yaml.MapSlice, key string) interface{} {
	for _, mapItem := range input {
		if mapItem.Key == key {
			return mapItem.Value
		}
	}

	return nil
}

func getPropertyString(input yaml.MapSlice, key string) string {
	value := getPropertyValue(input, key)
	if new, isOk := value.(string); isOk {
		return new
	}
	return ""
}

func postProcess(objects []RamlType) {
	objectMap := map[string]*RamlType{}

	for i, obj := range objects {
		objectMap[obj.Name] = &objects[i]
	}

	builtinTypes := []RamlType{}

	for i, obj := range builtinTypes {
		objectMap[obj.Name] = &builtinTypes[i]
	}

	for i := range objects {
		objects[i].resolve(objectMap)
	}

	for i := range objects {
		if objects[i].Discriminator != "" {
			objects[i].CodeName = "Abstract" + objects[i].CodeName
		} else {
			// TODO; check if 1 attribute and that is the discriminator
			if len(objects[i].Attributes) == 1 {
				for _, child := range objects[i].Children {
					if child.Discriminator != "" {
						objects[i].CodeName = "Abstract" + objects[i].CodeName
						break
					}
				}
			}
		}
	}

	// for _, obj := range objects {
	// 	fmt.Println(obj.Name)
	// 	fmt.Println(obj.Children)
	// }

}

func main() {
	fmt.Println("Generating code")

	f, err := ioutil.ReadFile("api.raml")

	t := yaml.MapSlice{}

	err = yaml.Unmarshal(f, &t)

	if err != nil {
		log.Fatalf("error: %v", err)
	}

	objects := parseYaml(t)
	postProcess(objects)
	generateCode(objects)
}
