# Bara - A GitOps tool for Kubernetes

[![Build Status](https://travis-ci.org/gary-lloyd-tessella/bara.svg?branch=master)](https://travis-ci.org/gary-lloyd-tessella/bara)

## Goal

To learn Go while building a micro service friendly Kubernetes deployment tool learning lessons from difficulties with Helm and Bara v1.

Bara should:

*	Support GitOps deployments by running as a service inside the cluster
*	Be easy to understand what is being applied
*	Be able to run in dry run mode
*	Notify on successful (or unsuccessful) deployment
*	Perform reconciliation

Input information:

*	Versioned deployment manifests - Questions - versioned how? helm charts? in another repo?
*	Deployment repo
*	Folder per environment
*	Flexible hierarchy for services
*	Leaf nodes are individual services if they contain the right files - metadata + values? or just values?