# Huck

[![Go](https://github.com/sdttttt/huck/actions/workflows/go.yml/badge.svg)](https://github.com/sdttttt/huck/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/sdttttt/huck/branch/master/graph/badge.svg?token=ZHKA97WxWR)](https://codecov.io/gh/sdttttt/huck)

Huck is a simple data collection server program, It's built on top of the Echo. that achieves the purpose of statistics through API calls. Huck is characterized by simple configuration and can be deployed without external network dependencies.

## Using

```yaml
# Statisticians used.
counter:

  #  Here is the route used
  - name: "测试员"
    path: "/week"
```

## Handlers

The following statistics are currently available:

- counter
