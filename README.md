# SPY CLI

**SpyCLI** is a CLI library for work with iac
projects that follows the **Blueprint** and **Thin Projects** with **[terragrunt](https://terragrunt.gruntwork.io/)**.

## Terragrunt project structure

Lets consider the following scenario:

We need provide an infrastrucutre. This infrastrucure need to have two environments:
- **environment-x**
- **environment-y**

And must be replicated three regions
- **region-a**
- **region-b**
- **region-c**

Each reagion must run two modules:
- **module-1**
- **module-2**

These requeriments results in a project schema similar to this:

![Infrastructure](assets/img/tg-project-schema.png)

And the files and folders structure must to be like this:

![Infrastructure](assets/img/tg-project-repeat.png)

Note that every region have the same modules and this files repeats each region folder. It's very bad deal with many files like that and the **Blueprint and Thin Projec Approach** intents to solve this problem.

## Blueprint and Thin Project approach

On this approach we remove each region modules from project and put that in a distributable and reusable **blueprint** that where this modules (per region or for any region) will be maintained and finaly the project will contains just essencial files (aka config files)

![Infrastructure](assets/img/bp-thin-project.png)

## SpyCli prject init

SpyCli helps developers to manage projects using scaffolds and automations. The most important is the command **init**. This commands creates the link between the project and the blueprint by copying or linkin files required

![Infrastructure](assets/img/tg-project-initialized.png)

---

## Commands Reference

### Access [SpyCLI commands reference](assets/docs/spycli.md) for more details