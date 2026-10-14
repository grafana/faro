# Adaptive sampling contract

The cross-language parts of Faro adaptive sampling. Every SDK implements this, the receiver validates against it, and the customer docs are derived from it.

- `rules.schema.json`: JSON Schema (2020-12) for the `faro.sampling.rules/v1` rules document a customer passes to the SDK. It encodes the field rules and the three rule kinds (exclusion, base rate, trigger). Two checks it can't express: rule ids must be unique in the document, and a rule without a trigger and without a scope makes every later untriggered rule unreachable (the SDK warns). Trigger types are not enumerated: an unknown type makes the rule inert with a warning, so a document written for a newer SDK degrades loudly rather than failing.
- `keys.json`: known-answer vectors for the sampling key. Each entry is a session id, its FNV-1a 32-bit hash and the final `keyU32`.
- `rolls.json`: the three roll inputs for a handful of session ids and rule ids, with their keys. One of the rule ids is `replay`, which must not share the replay roll.
- `rules-hash.json`: rule documents with their canonical form and `rulesHash`. Three are the same document spelled three ways and must hash the same; the rest each hit one place where JSON serializers tend to disagree.

A port of the key function, the roll inputs and the canonicalizer on any platform, and the receiver's recompute, must reproduce all three files bit for bit.

The verdict the SDK stamps on each payload is part of the OpenAPI spec: `components/schemas/sampling.yaml`, as `meta.sampling`.

## The key

```
h = FNV-1a 32-bit over the UTF-8 bytes of the input
      h = 0x811C9DC5; for each byte b: h = (h XOR b) * 0x01000193   mod 2^32
h = fmix32(h)                                   (MurmurHash3 finalizer)
      h = (h XOR (h >> 16)) * 0x85EBCA6B          mod 2^32
      h = (h XOR (h >> 13)) * 0xC2B2AE35          mod 2^32
      h =  h XOR (h >> 16)
keyU32 = h
```

Every multiply wraps mod 2^32 and every shift is a logical right shift. A roll passes when `keyU32(input) / 2^32 < rate`.

## The three rolls

Each roll hashes a different input, so the three are independent of each other and recomputable from the verdict:

```
base roll:     keyU32(sessionId)                                   / 2^32 < baseRate
replay roll:   keyU32(sessionId + "\0" + "replay" + "\0")           / 2^32 < replayRate
trigger roll:  keyU32(sessionId + "\0" + "trigger" + "\0" + ruleId) / 2^32 < triggerRate
```

`\0` is the NUL byte. Session ids match `^[a-zA-Z0-9_.@:-]+$` and rule ids `^[A-Za-z0-9_-]{1,64}$`, so neither can contain NUL, and no two of the three inputs can ever be the same string. In particular a rule with `id: "replay"` rolls on `sessionId\0trigger\0replay`, not on the replay roll's input.

## `rulesHash`

`rulesHash` identifies which rules document produced a verdict. It is FNV-1a 64-bit, as 16 lowercase hex characters, over the canonical form of the normalized document.

Normalization strips the two spellings that mean the same as absent: an explicit `"trigger": null` and an explicit `"scope": {}`. Nothing else is rewritten.

The canonical form is [RFC 8785 (JSON Canonicalization Scheme)](https://www.rfc-editor.org/rfc/rfc8785): object keys sorted by UTF-16 code units; no whitespace; strings escaped as ECMAScript `JSON.stringify` does (only `"`, `\` and U+0000 to U+001F are escaped, with `\b \f \n \r \t` short forms and lowercase `\u00xx` for the rest; no HTML escaping; non-ASCII written literally); numbers formatted as ECMAScript `Number.prototype.toString` (shortest round-trip digits, plain decimal for 1e-6 ≤ |x| < 1e21, otherwise `d.ddde±x` with no zero padding in the exponent). In JavaScript, `JSON.stringify` gives the right escaping and number formatting for scalars, and the default `Array.prototype.sort` gives the right key order, but the object must be written out key by key from the sorted list. Rebuilding an object and calling `JSON.stringify` on it does not work: JavaScript puts integer-like keys such as `"2"` and `"10"` in numeric order, while JCS sorts them as strings (`"10"` first). Other languages have to be told everything: Go's `encoding/json`, for example, escapes `<` as `\u003c`, formats `1e-7` as `1e-07`, and sorts keys by bytes. The documents in `rules-hash.json` cover each of those.

`rulesHash` is an audit identifier, not a security hash, and it is never an input to the key.
