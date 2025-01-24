#!/bin/bash

ver=$(git describe)
echo $ver > VERSION
