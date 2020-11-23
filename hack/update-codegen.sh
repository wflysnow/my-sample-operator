#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail
../vendor/k8s.io/code-generator/generate-groups.sh all \
my-sample-operator/pkg/generated \
my-sample-operator/pkg/apis \
samplecontroller:v1alpha1 \
--go-header-file $(pwd)/boilerplate.go.text \
--output-base $(pwd)/../../
