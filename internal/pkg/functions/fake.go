package functions

import "github.com/jaswdr/faker"

func fakeNS() *fakeNamespace {
	return &fakeNamespace{}
}

type fakeNamespace struct {
}

func (_ *fakeNamespace) Name() string {
	return faker.New().Person().Name()
}

func (_ *fakeNamespace) StreetAddress() string {
	return faker.New().Address().StreetAddress()
}

func (_ *fakeNamespace) City() string {
	return faker.New().Address().City()
}

func (_ *fakeNamespace) State() string {
	return faker.New().Address().State()
}
