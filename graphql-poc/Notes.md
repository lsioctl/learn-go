# Mutation example in graphqli:

```graphql
mutation CreateCharacter ($input: CharacterInput!) {
  upsertCharacter(input: $input) {
    name,
    id
  }
}
```

variables:

```json
{
  "input": {
   "name": "Kiara" 
  } 
}
```

Or:

```graphql
mutation {
  upsertCharacter(
   input: {
    name: "Kiara"
  }
  ) {
    name
    id
  }
}
````

# Query example

```graphql
query character($id: ID!) {
  character(id: $id) {
    id,
    name
  }
}
```

variables:

```json
{
  "id": 1
}
```

Or:

```graphql
query {
  character(id: 1) {
    name
    id
  }
}
```