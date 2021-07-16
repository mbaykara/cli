// Code generated by interfacer; DO NOT EDIT

package hcapi2

import (
	"context"
	"github.com/hetznercloud/hcloud-go/hcloud"
)

// SSHKeyClientBase is an interface generated for "github.com/hetznercloud/hcloud-go/hcloud.SSHKeyClient".
type SSHKeyClientBase interface {
	All(context.Context) ([]*hcloud.SSHKey, error)
	AllWithOpts(context.Context, hcloud.SSHKeyListOpts) ([]*hcloud.SSHKey, error)
	Create(context.Context, hcloud.SSHKeyCreateOpts) (*hcloud.SSHKey, *hcloud.Response, error)
	Delete(context.Context, *hcloud.SSHKey) (*hcloud.Response, error)
	Get(context.Context, string) (*hcloud.SSHKey, *hcloud.Response, error)
	GetByFingerprint(context.Context, string) (*hcloud.SSHKey, *hcloud.Response, error)
	GetByID(context.Context, int) (*hcloud.SSHKey, *hcloud.Response, error)
	GetByName(context.Context, string) (*hcloud.SSHKey, *hcloud.Response, error)
	List(context.Context, hcloud.SSHKeyListOpts) ([]*hcloud.SSHKey, *hcloud.Response, error)
	Update(context.Context, *hcloud.SSHKey, hcloud.SSHKeyUpdateOpts) (*hcloud.SSHKey, *hcloud.Response, error)
}
