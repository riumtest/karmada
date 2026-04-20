/*
Copyright 2024 The Karmada Authors.

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

package annotations

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetAnnotationValue retrieves the value of a specific annotation key from
// a Kubernetes object's annotations map. Returns the value and a boolean
// indicating whether the key was found.
//
// Example:
//
//	val, ok := GetAnnotationValue(obj.GetAnnotations(), "example.io/my-annotation")
//	if ok {
//		fmt.Println("Annotation value:", val)
//	}
func GetAnnotationValue(annotations map[string]string, key string) (string, bool) {
	if annotations == nil {
		return "", false
	}
	val, ok := annotations[key]
	return val, ok
}

// SetAnnotation sets a key-value pair in the annotations map of the provided
// ObjectMeta. If the annotations map is nil, it initializes it before setting.
func SetAnnotation(meta *metav1.ObjectMeta, key, value string) {
	if meta.Annotations == nil {
		meta.Annotations = make(map[string]string)
	}
	meta.Annotations[key] = value
}

// RemoveAnnotation removes the annotation with the given key from the
// ObjectMeta's annotations map. It is a no-op if the annotation does not exist.
func RemoveAnnotation(meta *metav1.ObjectMeta, key string) {
	if meta.Annotations == nil {
		return
	}
	delete(meta.Annotations, key)
}

// HasAnnotation returns true if the given annotations map contains the
// specified key, regardless of its value.
func HasAnnotation(annotations map[string]string, key string) bool {
	if annotations == nil {
		return false
	}
	_, ok := annotations[key]
	return ok
}

// MergeAnnotations merges src annotations into dst. If a key exists in both
// maps, the value from src takes precedence. Returns a new map without
// modifying the originals.
func MergeAnnotations(dst, src map[string]string) map[string]string {
	result := make(map[string]string, len(dst)+len(src))
	for k, v := range dst {
		result[k] = v
	}
	for k, v := range src {
		result[k] = v
	}
	return result
}

// ContainsAnnotations checks whether all provided key-value pairs exist in
// the annotations map with matching values.
// Note: if subset is empty, this always returns true (vacuously true).
func ContainsAnnotations(annotations map[string]string, subset map[string]string) bool {
	for k, v := range subset {
		if annotations[k] != v {
			return false
		}
	}
	return true
}
