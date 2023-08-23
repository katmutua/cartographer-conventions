//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020-2022 VMware Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by diegen. DO NOT EDIT.

package v1alpha1

import (
	json "encoding/json"
	fmtx "fmt"
	osx "os"
	reflectx "reflect"

	metav1 "dies.dev/apis/meta/v1"
	corev1 "k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	unstructured "k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	jsonpath "k8s.io/client-go/util/jsonpath"
	yaml "sigs.k8s.io/yaml"

	v1 "github.com/vmware-tanzu/cartographer-conventions/pkg/apis/thirdparty/cert-manager/v1"
)

var CertificateRequestBlank = (&CertificateRequestDie{}).DieFeed(v1.CertificateRequest{})

type CertificateRequestDie struct {
	metav1.FrozenObjectMeta
	mutable bool
	r       v1.CertificateRequest
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *CertificateRequestDie) DieImmutable(immutable bool) *CertificateRequestDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *CertificateRequestDie) DieFeed(r v1.CertificateRequest) *CertificateRequestDie {
	if d.mutable {
		d.FrozenObjectMeta = metav1.FreezeObjectMeta(r.ObjectMeta)
		d.r = r
		return d
	}
	return &CertificateRequestDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *CertificateRequestDie) DieFeedPtr(r *v1.CertificateRequest) *CertificateRequestDie {
	if r == nil {
		r = &v1.CertificateRequest{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *CertificateRequestDie) DieFeedJSON(j []byte) *CertificateRequestDie {
	r := v1.CertificateRequest{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *CertificateRequestDie) DieFeedYAML(y []byte) *CertificateRequestDie {
	r := v1.CertificateRequest{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *CertificateRequestDie) DieFeedYAMLFile(name string) *CertificateRequestDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *CertificateRequestDie) DieFeedRawExtension(raw runtime.RawExtension) *CertificateRequestDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *CertificateRequestDie) DieRelease() v1.CertificateRequest {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *CertificateRequestDie) DieReleasePtr() *v1.CertificateRequest {
	r := d.DieRelease()
	return &r
}

// DieReleaseUnstructured returns the resource managed by the die as an unstructured object. Panics on error.
func (d *CertificateRequestDie) DieReleaseUnstructured() *unstructured.Unstructured {
	r := d.DieReleasePtr()
	u, err := runtime.DefaultUnstructuredConverter.ToUnstructured(r)
	if err != nil {
		panic(err)
	}
	return &unstructured.Unstructured{
		Object: u,
	}
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *CertificateRequestDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *CertificateRequestDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *CertificateRequestDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *CertificateRequestDie) DieStamp(fn func(r *v1.CertificateRequest)) *CertificateRequestDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type or a pointer to that type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *CertificateRequestDie) DieStampAt(jp string, fn interface{}) *CertificateRequestDie {
	return d.DieStamp(func(r *v1.CertificateRequest) {
		if ni := reflectx.ValueOf(fn).Type().NumIn(); ni != 1 {
			panic(fmtx.Errorf("callback function must have 1 input parameters, found %d", ni))
		}
		if no := reflectx.ValueOf(fn).Type().NumOut(); no != 0 {
			panic(fmtx.Errorf("callback function must have 0 output parameters, found %d", no))
		}

		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			arg0t := reflectx.ValueOf(fn).Type().In(0)

			var args []reflectx.Value
			if cv.Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv}
			} else if cv.CanAddr() && cv.Addr().Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv.Addr()}
			} else {
				panic(fmtx.Errorf("callback function must accept value of type %q, found type %q", cv.Type(), arg0t))
			}

			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DieWith returns a new die after passing the current die to the callback function. The passed die is mutable.
func (d *CertificateRequestDie) DieWith(fns ...func(d *CertificateRequestDie)) *CertificateRequestDie {
	nd := CertificateRequestBlank.DieFeed(d.DieRelease()).DieImmutable(false)
	for _, fn := range fns {
		if fn != nil {
			fn(nd)
		}
	}
	return d.DieFeed(nd.DieRelease())
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *CertificateRequestDie) DeepCopy() *CertificateRequestDie {
	r := *d.r.DeepCopy()
	return &CertificateRequestDie{
		FrozenObjectMeta: metav1.FreezeObjectMeta(r.ObjectMeta),
		mutable:          d.mutable,
		r:                r,
	}
}

var _ runtime.Object = (*CertificateRequestDie)(nil)

func (d *CertificateRequestDie) DeepCopyObject() runtime.Object {
	return d.r.DeepCopy()
}

func (d *CertificateRequestDie) GetObjectKind() schema.ObjectKind {
	r := d.DieRelease()
	return r.GetObjectKind()
}

func (d *CertificateRequestDie) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.r)
}

func (d *CertificateRequestDie) UnmarshalJSON(b []byte) error {
	if d == CertificateRequestBlank {
		return fmtx.Errorf("cannot unmarshal into the blank die, create a copy first")
	}
	if !d.mutable {
		return fmtx.Errorf("cannot unmarshal into immutable dies, create a mutable version first")
	}
	r := &v1.CertificateRequest{}
	err := json.Unmarshal(b, r)
	*d = *d.DieFeed(*r)
	return err
}

// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
func (d *CertificateRequestDie) APIVersion(v string) *CertificateRequestDie {
	return d.DieStamp(func(r *v1.CertificateRequest) {
		r.APIVersion = v
	})
}

// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
func (d *CertificateRequestDie) Kind(v string) *CertificateRequestDie {
	return d.DieStamp(func(r *v1.CertificateRequest) {
		r.Kind = v
	})
}

// MetadataDie stamps the resource's ObjectMeta field with a mutable die.
func (d *CertificateRequestDie) MetadataDie(fn func(d *metav1.ObjectMetaDie)) *CertificateRequestDie {
	return d.DieStamp(func(r *v1.CertificateRequest) {
		d := metav1.ObjectMetaBlank.DieImmutable(false).DieFeed(r.ObjectMeta)
		fn(d)
		r.ObjectMeta = d.DieRelease()
	})
}

// SpecDie stamps the resource's spec field with a mutable die.
func (d *CertificateRequestDie) SpecDie(fn func(d *CertificateRequestSpecDie)) *CertificateRequestDie {
	return d.DieStamp(func(r *v1.CertificateRequest) {
		d := CertificateRequestSpecBlank.DieImmutable(false).DieFeed(r.Spec)
		fn(d)
		r.Spec = d.DieRelease()
	})
}

// StatusDie stamps the resource's status field with a mutable die.
func (d *CertificateRequestDie) StatusDie(fn func(d *CertificateRequestStatusDie)) *CertificateRequestDie {
	return d.DieStamp(func(r *v1.CertificateRequest) {
		d := CertificateRequestStatusBlank.DieImmutable(false).DieFeed(r.Status)
		fn(d)
		r.Status = d.DieRelease()
	})
}

// Desired state of the CertificateRequest resource.
func (d *CertificateRequestDie) Spec(v v1.CertificateRequestSpec) *CertificateRequestDie {
	return d.DieStamp(func(r *v1.CertificateRequest) {
		r.Spec = v
	})
}

// Status of the CertificateRequest. This is set and managed automatically.
func (d *CertificateRequestDie) Status(v v1.CertificateRequestStatus) *CertificateRequestDie {
	return d.DieStamp(func(r *v1.CertificateRequest) {
		r.Status = v
	})
}

var CertificateRequestSpecBlank = (&CertificateRequestSpecDie{}).DieFeed(v1.CertificateRequestSpec{})

type CertificateRequestSpecDie struct {
	mutable bool
	r       v1.CertificateRequestSpec
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *CertificateRequestSpecDie) DieImmutable(immutable bool) *CertificateRequestSpecDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *CertificateRequestSpecDie) DieFeed(r v1.CertificateRequestSpec) *CertificateRequestSpecDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &CertificateRequestSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *CertificateRequestSpecDie) DieFeedPtr(r *v1.CertificateRequestSpec) *CertificateRequestSpecDie {
	if r == nil {
		r = &v1.CertificateRequestSpec{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *CertificateRequestSpecDie) DieFeedJSON(j []byte) *CertificateRequestSpecDie {
	r := v1.CertificateRequestSpec{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *CertificateRequestSpecDie) DieFeedYAML(y []byte) *CertificateRequestSpecDie {
	r := v1.CertificateRequestSpec{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *CertificateRequestSpecDie) DieFeedYAMLFile(name string) *CertificateRequestSpecDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *CertificateRequestSpecDie) DieFeedRawExtension(raw runtime.RawExtension) *CertificateRequestSpecDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *CertificateRequestSpecDie) DieRelease() v1.CertificateRequestSpec {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *CertificateRequestSpecDie) DieReleasePtr() *v1.CertificateRequestSpec {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *CertificateRequestSpecDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *CertificateRequestSpecDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *CertificateRequestSpecDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *CertificateRequestSpecDie) DieStamp(fn func(r *v1.CertificateRequestSpec)) *CertificateRequestSpecDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type or a pointer to that type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *CertificateRequestSpecDie) DieStampAt(jp string, fn interface{}) *CertificateRequestSpecDie {
	return d.DieStamp(func(r *v1.CertificateRequestSpec) {
		if ni := reflectx.ValueOf(fn).Type().NumIn(); ni != 1 {
			panic(fmtx.Errorf("callback function must have 1 input parameters, found %d", ni))
		}
		if no := reflectx.ValueOf(fn).Type().NumOut(); no != 0 {
			panic(fmtx.Errorf("callback function must have 0 output parameters, found %d", no))
		}

		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			arg0t := reflectx.ValueOf(fn).Type().In(0)

			var args []reflectx.Value
			if cv.Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv}
			} else if cv.CanAddr() && cv.Addr().Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv.Addr()}
			} else {
				panic(fmtx.Errorf("callback function must accept value of type %q, found type %q", cv.Type(), arg0t))
			}

			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DieWith returns a new die after passing the current die to the callback function. The passed die is mutable.
func (d *CertificateRequestSpecDie) DieWith(fns ...func(d *CertificateRequestSpecDie)) *CertificateRequestSpecDie {
	nd := CertificateRequestSpecBlank.DieFeed(d.DieRelease()).DieImmutable(false)
	for _, fn := range fns {
		if fn != nil {
			fn(nd)
		}
	}
	return d.DieFeed(nd.DieRelease())
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *CertificateRequestSpecDie) DeepCopy() *CertificateRequestSpecDie {
	r := *d.r.DeepCopy()
	return &CertificateRequestSpecDie{
		mutable: d.mutable,
		r:       r,
	}
}

// The requested 'duration' (i.e. lifetime) of the Certificate. This option may be ignored/overridden by some issuer types.
func (d *CertificateRequestSpecDie) Duration(v *apismetav1.Duration) *CertificateRequestSpecDie {
	return d.DieStamp(func(r *v1.CertificateRequestSpec) {
		r.Duration = v
	})
}

// IssuerRef is a reference to the issuer for this CertificateRequest.  If the `kind` field is not set, or set to `Issuer`, an Issuer resource with the given name in the same namespace as the CertificateRequest will be used.  If the `kind` field is set to `ClusterIssuer`, a ClusterIssuer with the provided name will be used. The `name` field in this stanza is required at all times. The group field refers to the API group of the issuer which defaults to `cert-manager.io` if empty.
func (d *CertificateRequestSpecDie) IssuerRef(v corev1.ObjectReference) *CertificateRequestSpecDie {
	return d.DieStamp(func(r *v1.CertificateRequestSpec) {
		r.IssuerRef = v
	})
}

// The PEM-encoded x509 certificate signing request to be submitted to the CA for signing.
func (d *CertificateRequestSpecDie) Request(v []byte) *CertificateRequestSpecDie {
	return d.DieStamp(func(r *v1.CertificateRequestSpec) {
		r.Request = v
	})
}

// IsCA will request to mark the certificate as valid for certificate signing when submitting to the issuer. This will automatically add the `cert sign` usage to the list of `usages`.
func (d *CertificateRequestSpecDie) IsCA(v bool) *CertificateRequestSpecDie {
	return d.DieStamp(func(r *v1.CertificateRequestSpec) {
		r.IsCA = v
	})
}

// Usages is the set of x509 usages that are requested for the certificate. If usages are set they SHOULD be encoded inside the CSR spec Defaults to `digital signature` and `key encipherment` if not specified.
func (d *CertificateRequestSpecDie) Usages(v ...v1.KeyUsage) *CertificateRequestSpecDie {
	return d.DieStamp(func(r *v1.CertificateRequestSpec) {
		r.Usages = v
	})
}

// Username contains the name of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.
func (d *CertificateRequestSpecDie) Username(v string) *CertificateRequestSpecDie {
	return d.DieStamp(func(r *v1.CertificateRequestSpec) {
		r.Username = v
	})
}

// UID contains the uid of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.
func (d *CertificateRequestSpecDie) UID(v string) *CertificateRequestSpecDie {
	return d.DieStamp(func(r *v1.CertificateRequestSpec) {
		r.UID = v
	})
}

// Groups contains group membership of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.
func (d *CertificateRequestSpecDie) Groups(v ...string) *CertificateRequestSpecDie {
	return d.DieStamp(func(r *v1.CertificateRequestSpec) {
		r.Groups = v
	})
}

// Extra contains extra attributes of the user that created the CertificateRequest. Populated by the cert-manager webhook on creation and immutable.
func (d *CertificateRequestSpecDie) Extra(v map[string][]string) *CertificateRequestSpecDie {
	return d.DieStamp(func(r *v1.CertificateRequestSpec) {
		r.Extra = v
	})
}

var CertificateRequestStatusBlank = (&CertificateRequestStatusDie{}).DieFeed(v1.CertificateRequestStatus{})

type CertificateRequestStatusDie struct {
	mutable bool
	r       v1.CertificateRequestStatus
}

// DieImmutable returns a new die for the current die's state that is either mutable (`false`) or immutable (`true`).
func (d *CertificateRequestStatusDie) DieImmutable(immutable bool) *CertificateRequestStatusDie {
	if d.mutable == !immutable {
		return d
	}
	d = d.DeepCopy()
	d.mutable = !immutable
	return d
}

// DieFeed returns a new die with the provided resource.
func (d *CertificateRequestStatusDie) DieFeed(r v1.CertificateRequestStatus) *CertificateRequestStatusDie {
	if d.mutable {
		d.r = r
		return d
	}
	return &CertificateRequestStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

// DieFeedPtr returns a new die with the provided resource pointer. If the resource is nil, the empty value is used instead.
func (d *CertificateRequestStatusDie) DieFeedPtr(r *v1.CertificateRequestStatus) *CertificateRequestStatusDie {
	if r == nil {
		r = &v1.CertificateRequestStatus{}
	}
	return d.DieFeed(*r)
}

// DieFeedJSON returns a new die with the provided JSON. Panics on error.
func (d *CertificateRequestStatusDie) DieFeedJSON(j []byte) *CertificateRequestStatusDie {
	r := v1.CertificateRequestStatus{}
	if err := json.Unmarshal(j, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAML returns a new die with the provided YAML. Panics on error.
func (d *CertificateRequestStatusDie) DieFeedYAML(y []byte) *CertificateRequestStatusDie {
	r := v1.CertificateRequestStatus{}
	if err := yaml.Unmarshal(y, &r); err != nil {
		panic(err)
	}
	return d.DieFeed(r)
}

// DieFeedYAMLFile returns a new die loading YAML from a file path. Panics on error.
func (d *CertificateRequestStatusDie) DieFeedYAMLFile(name string) *CertificateRequestStatusDie {
	y, err := osx.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return d.DieFeedYAML(y)
}

// DieFeedRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *CertificateRequestStatusDie) DieFeedRawExtension(raw runtime.RawExtension) *CertificateRequestStatusDie {
	j, err := json.Marshal(raw)
	if err != nil {
		panic(err)
	}
	return d.DieFeedJSON(j)
}

// DieRelease returns the resource managed by the die.
func (d *CertificateRequestStatusDie) DieRelease() v1.CertificateRequestStatus {
	if d.mutable {
		return d.r
	}
	return *d.r.DeepCopy()
}

// DieReleasePtr returns a pointer to the resource managed by the die.
func (d *CertificateRequestStatusDie) DieReleasePtr() *v1.CertificateRequestStatus {
	r := d.DieRelease()
	return &r
}

// DieReleaseJSON returns the resource managed by the die as JSON. Panics on error.
func (d *CertificateRequestStatusDie) DieReleaseJSON() []byte {
	r := d.DieReleasePtr()
	j, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}
	return j
}

// DieReleaseYAML returns the resource managed by the die as YAML. Panics on error.
func (d *CertificateRequestStatusDie) DieReleaseYAML() []byte {
	r := d.DieReleasePtr()
	y, err := yaml.Marshal(r)
	if err != nil {
		panic(err)
	}
	return y
}

// DieReleaseRawExtension returns the resource managed by the die as an raw extension. Panics on error.
func (d *CertificateRequestStatusDie) DieReleaseRawExtension() runtime.RawExtension {
	j := d.DieReleaseJSON()
	raw := runtime.RawExtension{}
	if err := json.Unmarshal(j, &raw); err != nil {
		panic(err)
	}
	return raw
}

// DieStamp returns a new die with the resource passed to the callback function. The resource is mutable.
func (d *CertificateRequestStatusDie) DieStamp(fn func(r *v1.CertificateRequestStatus)) *CertificateRequestStatusDie {
	r := d.DieRelease()
	fn(&r)
	return d.DieFeed(r)
}

// Experimental: DieStampAt uses a JSON path (http://goessner.net/articles/JsonPath/) expression to stamp portions of the resource. The callback is invoked with each JSON path match. Panics if the callback function does not accept a single argument of the same type or a pointer to that type as found on the resource at the target location.
//
// Future iterations will improve type coercion from the resource to the callback argument.
func (d *CertificateRequestStatusDie) DieStampAt(jp string, fn interface{}) *CertificateRequestStatusDie {
	return d.DieStamp(func(r *v1.CertificateRequestStatus) {
		if ni := reflectx.ValueOf(fn).Type().NumIn(); ni != 1 {
			panic(fmtx.Errorf("callback function must have 1 input parameters, found %d", ni))
		}
		if no := reflectx.ValueOf(fn).Type().NumOut(); no != 0 {
			panic(fmtx.Errorf("callback function must have 0 output parameters, found %d", no))
		}

		cp := jsonpath.New("")
		if err := cp.Parse(fmtx.Sprintf("{%s}", jp)); err != nil {
			panic(err)
		}
		cr, err := cp.FindResults(r)
		if err != nil {
			// errors are expected if a path is not found
			return
		}
		for _, cv := range cr[0] {
			arg0t := reflectx.ValueOf(fn).Type().In(0)

			var args []reflectx.Value
			if cv.Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv}
			} else if cv.CanAddr() && cv.Addr().Type().AssignableTo(arg0t) {
				args = []reflectx.Value{cv.Addr()}
			} else {
				panic(fmtx.Errorf("callback function must accept value of type %q, found type %q", cv.Type(), arg0t))
			}

			reflectx.ValueOf(fn).Call(args)
		}
	})
}

// DieWith returns a new die after passing the current die to the callback function. The passed die is mutable.
func (d *CertificateRequestStatusDie) DieWith(fns ...func(d *CertificateRequestStatusDie)) *CertificateRequestStatusDie {
	nd := CertificateRequestStatusBlank.DieFeed(d.DieRelease()).DieImmutable(false)
	for _, fn := range fns {
		if fn != nil {
			fn(nd)
		}
	}
	return d.DieFeed(nd.DieRelease())
}

// DeepCopy returns a new die with equivalent state. Useful for snapshotting a mutable die.
func (d *CertificateRequestStatusDie) DeepCopy() *CertificateRequestStatusDie {
	r := *d.r.DeepCopy()
	return &CertificateRequestStatusDie{
		mutable: d.mutable,
		r:       r,
	}
}

// List of status conditions to indicate the status of a CertificateRequest. Known condition types are `Ready` and `InvalidRequest`.
func (d *CertificateRequestStatusDie) Conditions(v ...v1.CertificateRequestCondition) *CertificateRequestStatusDie {
	return d.DieStamp(func(r *v1.CertificateRequestStatus) {
		r.Conditions = v
	})
}

// The PEM encoded x509 certificate resulting from the certificate signing request. If not set, the CertificateRequest has either not been completed or has failed. More information on failure can be found by checking the `conditions` field.
func (d *CertificateRequestStatusDie) Certificate(v []byte) *CertificateRequestStatusDie {
	return d.DieStamp(func(r *v1.CertificateRequestStatus) {
		r.Certificate = v
	})
}

// The PEM encoded x509 certificate of the signer, also known as the CA (Certificate Authority). This is set on a best-effort basis by different issuers. If not set, the CA is assumed to be unknown/not available.
func (d *CertificateRequestStatusDie) CA(v []byte) *CertificateRequestStatusDie {
	return d.DieStamp(func(r *v1.CertificateRequestStatus) {
		r.CA = v
	})
}

// FailureTime stores the time that this CertificateRequest failed. This is used to influence garbage collection and back-off.
func (d *CertificateRequestStatusDie) FailureTime(v *apismetav1.Time) *CertificateRequestStatusDie {
	return d.DieStamp(func(r *v1.CertificateRequestStatus) {
		r.FailureTime = v
	})
}
