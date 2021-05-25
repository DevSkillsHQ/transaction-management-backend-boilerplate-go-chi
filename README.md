# Transaction Management Backend Boilerplate

A Go-based boilerplate for the Transaction Management Backend coding challenge from [DevSkills](http://devskills.co/).

## Contents

- [app/server.go](app/server.go) - a Golang service with predefined dummy endpoints for the [Transaction Management API](https://infra.devskills.app/transaction-management/api/3.0.0).
- [db/db.go](db/db.go) - a SQLite setup that re-creates a SQLite DB from scratch and connects to it.
- [API test suite](https://github.com/DevSkillsHQ/backend-boilerplate-golang/blob/main/cypress/integration/backend.spec.js) - a dummy Cypress test suite.
- [Pipeline](https://github.com/DevSkillsHQ/backend-boilerplate-golang/blob/main/.github/workflows/tests.yml) - a test Runner that executes the Cypress tests on push to a branch other than `master`/`main`.

## Tech Stack

- [Golang](https://golang.org/)
- [Chi](https://github.com/go-chi/chi)
- [Cypress](https://www.cypress.io/)
- [GitHub Actions](https://github.com/features/actions)

---

Made by [DevSkills](https://devskills.co).

Did you find this repo useful? **Give us a shout on [Twitter](https://twitter.com/DevSkillsHQ) / [LinkedIn](https://www.linkedin.com/company/devskills)**.
