#!/bin/bash -eux

pushd dp-nlp-hub
  make build
  cp build/dp-nlp-hub Dockerfile.concourse ../build
popd
