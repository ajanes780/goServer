# USER
- id : string
- name : string
- email : string
- password : string
- comments : Comment[]

# Comments
- id : string
- content : string
- user : UserID
- article : ArticleID

# Author

- id : string
- name : string
- email : string
- bio : string
- articles : []ArticleID

# Article

- author : Author
- title : string
- content : string
- summary : string
- published : boolean
- likes : number
