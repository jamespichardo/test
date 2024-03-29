# QuakBot Trojan

<kbd>[UNIT:ANTIVIRUS](https://docs.preludesecurity.com/docs/security-policy#antivirus)</kbd>
<kbd>[ALERT:-](#the-url)</kbd>

This test drops a defanged QuakBot OneNote file. OneNote allows attackers to embed malicious files such as .hta and executables. Malicious OneNote files require users to click on the embedded file for execution. While there is a pop-up warning message when clicking on the object, most users will click through and allow the object to be executed, in turn infecting their system. QuakBot is a banking trojan primarily used to steal financial data, including browser information, keystrokes, and credentials. Once QakBot has successfully infected an environment, the malware installs a backdoor allowing the threat actor to drop additional malware and has been commonly used by larger ransomware groups such as BlackBasta. This test will monitor if any endpoint defense quarantines the malware.

## How

> Safety: the malware used has been defanged, so even if run, it will immediately exit.

The malware file is extracted from the test and written to a user-owned directory. The test then waits briefly before running a few checks to determine if it was responded to (not just detected) by any defensive products on the endpoint.

Example Output:
```bash
[+] Extracting file for quarantine test
[+] Pausing for 3 seconds to gauge defensive reaction
[-] Malicious file was not caught
```

## Resolution

If this test fails:

* Ensure you have an antivirus program installed and running.
* If using an EDR, make sure the antivirus capability is enabled and turned up, appropriately
