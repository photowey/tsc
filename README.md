# `tsc`
> <font style="color:red">T</font>ype<font style="color:red">S</font>cript model <font style="color:red">C</font>onverter
>
> A project that converts `json` data structures into `typescript` data models


English | [中文](./README_CN.md)



## 1.`Usage`

- `tsc -d '${json.string.content}'`

  - ```shell
    $ tsc -d '"{\"name|required\":\"photowey\",\"age|readonly\":18}"'
    
    # Use single quotes('') to wrap json string content
    ```

- `tsc --data '${json.string.content}'`

  - ```shell
    $ tsc --data '"{\"name|required\":\"photowey\",\"age|readonly\":18}"'
    ```

- `tsc in ${json.file}`

  - > Not supported yet: `Linux` input redirection, replace **`<`** with the **`in`** subcommand

  - ```shell
  $ tsc in example.json
    ```
  
  - ```shell
  $ tsc in example1.json example2.json
    ```

  - ```shell
    $ tsc in example1.json ... exampleN.json
    ```




## 2.`Output`:

### 2.1.`json file`

```json
{
  "name|required": "photowey",
  "age|readonly": 18,
  "balance": 10.24,
  "boy": true,
  "address": "chongqing",
  "hobby": [
    {
      "name": "badminton",
      "description": "badminton"
    }
  ],
  "university": {
    "name": "xxx",
    "address": "ertang"
  }
}
```

```shell
# Notes:
# We can decorate the `key` of the `json` data structure with `required` and `readonly`.
# Thus, control over `TS` data model properties is achieved
```

### 2.2.`typescript model`

> the default top data-model name is "**`DataModel`**"
>
> > Since `GO` converts `JSON` content into `map[string]any` when processing it, it is also an unordered data structure.
> >
> > > This will lead to: The order of the attributes of the data model output is not the same each time.
> >
> > supported data types:
> >
> > - `string`
> > - `number`
> > - `boolean`

```typescript
// ---------------------------------------------------------------- example.json
export interface DataModel {
    readonly age?: number // 18
    balance?: number // 10.24
    boy?: boolean // true
    address?: string // chongqing
    Hobby?: Hobby[] // Hobby
    University?: University // University
    name: string // photowey
}

export interface Hobby {
    name?: string // badminton
    description?: string // badminton
}

export interface University {
    address?: string // ertang
    name?: string // xxx
}
```