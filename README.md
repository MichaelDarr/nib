# nib
**Generate SQL migrations from GORM models**

Nib is a database migration CLI marrying [GORM](https://github.com/go-gorm/gorm) models to [migrate](https://github.com/golang-migrate/migrate) migrations. [Alembic](https://alembic.sqlalchemy.org), SQLAlchemy's database migration tool, inspired this project. 

## Table of Contents
* [Installation](#installation)

## Installation

**Prerequisites:**
* git
* go (tested on 1.13.7)
* make

```sh
$ git clone git@github.com:MichaelDarr/nib.git
$ cd nib
$ make # or, if ahab is installed, build in container with `make build-ahab`
$ sudo make install
```
