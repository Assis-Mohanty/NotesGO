Build a secure REST API in Golang (Fiber) for user authentication and personal notes management. Only authenticated users can access and manage their own notes.

TODO :/register //TODO -

TODO :/login   //TODO -

TODO :POST /notes – create a note (title, content) //TODO -

TODO GET /notes – list all notes of logged-in user //TODO - 

TODO GET /notes/:id, PUT, DELETE – only for user's own notes //TODO - 

TODO Notes pagination or search by query param  //TODO-

TODO CLI to seed dummy users and notes  //TODO - 



**Requirements:**

* **Auth Routes:**

  * POST /register – name, email, password (hashed with bcrypt)
  * POST /login – returns JWT

* **Note Routes (JWT protected):**

  * POST /notes – create a note (title, content)
  * GET /notes – list all notes of logged-in user
  * GET /notes/:id, PUT, DELETE – only for user's own notes

* **Tech Stack:**

  * Golang + Fiber
  * MySQL (use GORM)
  * JWT for auth (custom middleware)
  * Docker: app + MySQL setup with docker-compose
  * Env config using .env file

* **Folder Structure:**

  * /models, /routes, /handlers, /middleware, /utils

**Constraints:**

* No boilerplate templates or full-stack generators
* Must write JWT + bcrypt logic manually
* Return proper status codes & validation (empty fields, unauthorized access)

**Bonus:**

* Notes pagination or search by query param
* CLI to seed dummy users and notes 