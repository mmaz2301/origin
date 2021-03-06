package v1

import (
	v1 "github.com/openshift/origin/pkg/image/apis/image/v1"
	scheme "github.com/openshift/origin/pkg/image/generated/clientset/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ImagesGetter has a method to return a ImageResourceInterface.
// A group's client should implement this interface.
type ImagesGetter interface {
	Images() ImageResourceInterface
}

// ImageResourceInterface has methods to work with ImageResource resources.
type ImageResourceInterface interface {
	Create(*v1.Image) (*v1.Image, error)
	Update(*v1.Image) (*v1.Image, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.Image, error)
	List(opts meta_v1.ListOptions) (*v1.ImageList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Image, err error)
	ImageResourceExpansion
}

// images implements ImageResourceInterface
type images struct {
	client rest.Interface
}

// newImages returns a Images
func newImages(c *ImageV1Client) *images {
	return &images{
		client: c.RESTClient(),
	}
}

// Create takes the representation of a imageResource and creates it.  Returns the server's representation of the imageResource, and an error, if there is any.
func (c *images) Create(imageResource *v1.Image) (result *v1.Image, err error) {
	result = &v1.Image{}
	err = c.client.Post().
		Resource("images").
		Body(imageResource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a imageResource and updates it. Returns the server's representation of the imageResource, and an error, if there is any.
func (c *images) Update(imageResource *v1.Image) (result *v1.Image, err error) {
	result = &v1.Image{}
	err = c.client.Put().
		Resource("images").
		Name(imageResource.Name).
		Body(imageResource).
		Do().
		Into(result)
	return
}

// Delete takes name of the imageResource and deletes it. Returns an error if one occurs.
func (c *images) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("images").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *images) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Resource("images").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the imageResource, and returns the corresponding imageResource object, and an error if there is any.
func (c *images) Get(name string, options meta_v1.GetOptions) (result *v1.Image, err error) {
	result = &v1.Image{}
	err = c.client.Get().
		Resource("images").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Images that match those selectors.
func (c *images) List(opts meta_v1.ListOptions) (result *v1.ImageList, err error) {
	result = &v1.ImageList{}
	err = c.client.Get().
		Resource("images").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested images.
func (c *images) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("images").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched imageResource.
func (c *images) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Image, err error) {
	result = &v1.Image{}
	err = c.client.Patch(pt).
		Resource("images").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
