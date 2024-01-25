# PixiePress

TODO: Write documentations

The database shape is as follows:

1. Articles
    - Title
    - Author (should relate to the "Users" entity)
    - Publication Date
    - Summary
    - Tags (should be synonymous with keywords)
    - Cover Image
    - Category (should relate to the "Categories" entity)
2. Categories
    - Name
    - Articles (should relate to a list of "Articles" entity)
    - Description
    - MetaImage (for SEO purposes)
3. Users
    - Name
    - Profile Picture
    - Biography
    - Articles (should to a list of the "Articles" entity)
    - Role (should relate to the "Roles" entity)
4. Roles
    - Name
    - Users (should relate to the "Users" entity)

```mermaid
erDiagram
    ARTICLE {
        int articleId PK
        string title
        int authorId FK
        int publicationDate
        string summary
        string[] tags
        string coverImage
        string categoryName FK
    }
    AUTHOR {
        int authorId PK
        string username
        string profilePicture
        string biogpraphy
        int[] articleId FK
        string roleName FK
        string email
        string password
    }
    CATEGORY {
        string categoryName PK
        int[] articleId FK
        string description
        string metaImage
    }
    ROLE {
        string roleName PK
        int[] authorId FK
    }

    ARTICLE ||--|| CATEGORY : has
    CATEGORY ||--|{ ARTICLE : contains

    ARTICLE ||--|| AUTHOR : has
    AUTHOR ||--|{ ARTICLE : has

    AUTHOR ||--|| ROLE : has
    ROLE ||--|{ AUTHOR : contains
````
