# Secretify cli

## Create

```bash
secretify create -u https://beta.secretify.io/api/v1 -s foo
# Output: 'GdqitgHPCp4prwdWhH#GfzGXBU8f97BiRRTgVeWip6kVP'
```

## Reveal

```bash
secretify reveal -u https://beta.secretify.io/api/v1 -i 'GdqitgHPCp4prwdWhH#GfzGXBU8f97BiRRTgVeWip6kVP'
# Output: 'foo'
```

**Note**: Currently you can not use the identifier on the UI and vice versa.
