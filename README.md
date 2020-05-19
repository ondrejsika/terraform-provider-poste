> This repository is in **Work in Progress** state. If you need something, create an [issue](https://github.com/ondrejsika/terraform-provider-poste/issues/new)

# terraform-provider-poste

    2019 Ondrej Sika <ondrej@ondrejsika.com>
    https://github.com/ondrejsika/terraform-provider-poste

## Example Usage

```terraform
provider "poste" {
  origin = "http://localhost/admin/api"
  username = "admin@example.com"
  password = "asdfasdf"
}

resource "poste_domain" "foo_com" {
  name = "foo.com"
}

resource "poste_domain" "bar_com" {
  name = "bar.com"
}

resource "poste_box" "foo_foo_com" {
  email = "foo@${poste_domain.foo_com.name}"
  password = "asdfasdf1"
}

resource "poste_box" "noreply_foo_com" {
  email = "noreply@${poste_domain.foo_com.name}"
  password = "asdfasdf1"
}

resource "poste_box" "bar_foobar_com" {
  email = "bar@${poste_domain.foobar_com.name}"
  password = "asdfasdf1"
}
```

## Change Log

- v0.2.0 - Add `poste_box` resource (only add/remove box with plaintext password)

- v0.1.0 - First version of provider with `poste_domain` resource (only add/remove domain)
