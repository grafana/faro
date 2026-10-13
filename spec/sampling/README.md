# Adaptive sampling contract

The cross-language parts of Faro adaptive sampling. Every SDK implements this, the receiver validates against it, and the customer docs are derived from it.

- `rules.schema.json`: JSON Schema (2020-12) for the `faro.sampling.rules/v1` rules document a customer passes to the SDK. It encodes the field rules and the three rule kinds (exclusion, base rate, trigger). Two checks it can't express: rule ids must be unique in the document, and a rule without a trigger and without a scope makes every later untriggered rule unreachable (the SDK warns). Trigger types are not enumerated: an unknown type makes the rule inert with a warning, so a document written for a newer SDK degrades loudly rather than failing.
- `keys.json`: known-answer vectors for the sampling key. Each entry is a session id, its FNV-1a 32-bit hash and the final `keyU32`. A port of the key function on any platform, and the receiver's recompute, must reproduce these bit for bit.

The verdict the SDK stamps on each payload is part of the OpenAPI spec: `components/schemas/sampling.yaml`, as `meta.sampling`.

## The key

```
h = FNV-1a 32-bit over the UTF-8 bytes of the session id
      h = 0x811C9DC5; for each byte b: h = (h XOR b) * 0x01000193   mod 2^32
h = fmix32(h)                                   (MurmurHash3 finalizer)
      h = (h XOR (h >> 16)) * 0x85EBCA6B          mod 2^32
      h = (h XOR (h >> 13)) * 0xC2B2AE35          mod 2^32
      h =  h XOR (h >> 16)
keyU32 = h
```

Every multiply wraps mod 2^32 and every shift is a logical right shift. The base roll is `keyU32(sessionId) / 2^32 < baseRate`. The replay roll uses `keyU32(sessionId + "\0" + "replay")`, and a trigger rule's `triggerRate` roll uses `keyU32(sessionId + "\0" + ruleId)`, so each roll is independent of the others and recomputable from the verdict.

`rulesHash` is FNV-1a 64-bit, as 16 hex characters, over the canonical JSON (keys sorted, no whitespace) of the rules document after normalisation: an explicit `"trigger": null` and an explicit `"scope": {}` are stripped first, since they mean the same as absent. It identifies which document produced a verdict. It is never an input to the key.
