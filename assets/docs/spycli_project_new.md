## spycli project new

Create new project

### Synopsis

Creates a new project with local or remote reference for blueprint and components

Examples:

Create a project that:

- Is called "My Project"
- Uses blueprint bp-aws-nearform and stack simple-web-app locally
- Have two environments: develop and production
- Have two regions: us-east-1 and us-west-1

spycli project new -n "My Project" -b bp-aws-nearform -l -s simple-web-app -r us-east-1 -e develop -e production

The same project but using remove blueprint:

spycli project new -n "My Project" -b git@github.com:nearform/bp-aws-nearform.git -s simple-web-app -r us-east-1 -e develop -e production

The same project but using remote blueprint and remote state in terraform:

spycli project new -n "My Project" -b git@github.com:nearform/bp-aws-nearform.git -s simple-web-app -r us-east-1 -e develop -e production -t -u my-bucket -v us-east-1



```
spycli project new [flags]
```

### Options

```
  -b, --blueprint string              Blueprint
  -d, --directory string              Base directory where the files will be writen (default ".")
  -e, --environment strings           Pass a list of environments (default [dev])
  -h, --help                          help for new
  -l, --local                         Local blueprint
  -n, --name string                   Element name (ex: my-project or my-blueprint)
  -p, --platform string               Plataform or service (aws|azure) (default "aws")
  -r, --region strings                Pass a list of environments (default [us-east-1])
  -u, --remote-bucket string          Stack name
  -v, --remote-bucket-region string   Stack name
  -t, --remote-state                  Use remote state
  -s, --stack string                  Stack name
```

### SEE ALSO

* [spycli project](spycli_project.md)	 - Manipulate iac projects

###### Auto generated by spf13/cobra on 25-Jan-2022
