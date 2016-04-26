package do

import (
	"github.com/asarturas/dodin/cloud"
	"github.com/digitalocean/godo"
	"golang.org/x/oauth2"
)

func Provider(apiToken string) DigitalOceanProvider {
	return DigitalOceanProvider{
		apiToken: apiToken,
	}
}

type DigitalOceanProvider struct {
	apiToken string
}

type TokenSource struct {
	AccessToken string
}

func (t *TokenSource) Token() (*oauth2.Token, error) {
	token := &oauth2.Token{
		AccessToken: t.AccessToken,
	}
	return token, nil
}

func (provider DigitalOceanProvider) GetMachines() []cloud.Machine {
	tokenSource := &TokenSource{
		AccessToken: provider.apiToken,
	}
	oauthClient := oauth2.NewClient(oauth2.NoContext, tokenSource)
	client := godo.NewClient(oauthClient)
	opt := &godo.ListOptions{
		Page:    1,
		PerPage: 200,
	}
	droplets, _, _ := client.Droplets.List(opt)
	machines := []cloud.Machine{}
	for _, droplet := range droplets {
		ip, _ := droplet.PublicIPv4()
		machines = append(machines, cloud.GetCloudMachine(droplet.Name, ip))
	}
	return machines
}
