# Kontext

Think of this as a different direnv that does not put anything into env and where every value is in a separate file.

## Installation

```[bash]
go install -ldflags="-s -w"
```

## Usage

```[bash]
# this will fail b/c there is no context file yet
kontext get project

#create a context file with a uuid in it
uuidgen > .project.kontext

# this will show the full path of the context file
kontext get project

# this will output the generated context uuid
# this will also work when you are one directory deeper
cat $(kontext project)
```

### Use a context with a JSON store

```[bash]
jq -n --arg date "$(date -u +"%Y-%m-%dT%H:%M:%SZ")" \
      --arg uuid "$(uuidgen)" \
      '{created: $date, id: $id}' > .project.kontext

jq -r '.uuid' "$(kontext get project)"
jq -r '.created' "$(kontext get project)"
```

### Use context as a taskwarrior project

You can use kontext to

```[bash]
uuidgen > .project.kontext
task add project:$(kontext get project) "this is a new task for THIS project"
task add project:$(kontext get project) "this is another new task for THIS project"

# list all tasks for this project
task project:$(kontext get project)
```

Get wild! <3
