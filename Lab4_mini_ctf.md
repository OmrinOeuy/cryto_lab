# Mini-CTF (Cat Leak)

**Objective:** Identify the leaked information, analyze the encoded URL, understand the hash-related clues, and determine the correct CTF flag.

## Step 1 - Analyze the Google search URL

**Example:** `%62%36%63%63%62%34`

## Step 2 - Decoding the URL

By using this Python function:

```python
from urllib.parse import unquote

encoded = "%62%36%63%63%62%34%65%63%65%35%34%35%34%64%63%61%65%35%31%37%37%38%62%33%65%32%33%39%65%62%63%32"
decoded = unquote(encoded)

print("Decoded value:", decoded)
```

**Output:**
```text
Decoded value: b6ccb4ece5454dcae51778b3e239ebc2
```

## Step 3 - Analyze the Hash

* The decoded value is a 32-character hex string.
* The challenge mentions bcrypt, but the supplied value alone does not establish which algorithm generated it.
* A leaked hash is not necessarily the original plaintext password.

## Step 4 - Analyze the Flag 

`^cryptoCTF\{(?:\x6d\x65\x6f\x77){2}\}$`

| Regex | Meaning |
| :--- | :--- |
| `\x6d` | m |
| `\x65` | e |
| `\x6f` | o |
| `\x77` | w |
| `{2}` | Repeat twice |

**Final Flag:**
```text
cryptoCTF{meowmeow}
```