# Module name: `lba-enrichment`

> Localize preys using the localization by association (LBA) technique

| parameters | definition | default |
|------------|------------|---------|
| database | fasta database for mapping gene names to IDs | |
| fdr | FDR threshold for defining significant preys | 0.01 |
| minBaits | minimum baits a prey must be seen with to be included in the analysis | 1 |
| minFC | minimum fold change to include a prey in an association list (exclusive) | 1 |
| namespace | GO namespace to use, one of BP, CC, MF | CC |
| outFile | output file name for prey localization data | lba-enrichment.txt |
| preyLimit | the maximum number of associated preys to use for defining prey profile | 100 |
| saintFile | SAINT txt file | |

## Example file formats

### database
```
>NP_001263222.1|gn|NISCH:11188| nischarin isoform 2 [Homo sapiens]
MATARTFGPEREAEPAKEARVVGSELVDTYTVYIIQVTDGSHEWTVKHRYSDFHDLHEKLVAERKIDKNL
LPPKKIIGKNSRSLVEKREKDLEVYLQKLLAAFPGVTPRVLAHFLHFHFYEINGITAALAEELFEKGEQL
LGAGEVFAIGPLQLYAVTEQLQQGKPTCASGDAKTDLGHILDFTCRLKYLKVSGTEGPFGTSNIQEQLLP
FDLSIFKSLHQVEISHCDAKHIRGLVASKPTLATLSVRFSATSMKEVLVPEASEFDEWEPEGTTLEGPVT
AVIPTWQALTTLDLSHNSISEIDESVKLIPKIEFLDLSHNGLLVVDNLQHLYNLVHLDLSYNKLSSLEGL
HTKLGNIKTLNLAGNLLESLSGLHKLYSLVNLDLRDNRIEQMEEVRSIGSLPCLEHVSLLNNPLSIIPDY
RTKVLAQFGERASEVCLDDTVTTEKELDTVEVLKAIQKAKEVKSKLSNPEKKGGEDSRLSAAPCIRPSSS
PPTVAPASASLPQPILSNQGNRVCILLLVEPHSPAWAPWLGWGWGRGASTCFQQGQTQGGQCLLQAGPRG
GTHGRGAWPDASCCLLGEDSQLL
>NP_001263218.1|gn|PRKAR1A:5573| cAMP-dependent protein kinase type I-alpha regulatory subunit isoform a [Homo sapiens]
MESGSTAASEEARSLRECELYVQKHNIQALLKDSIVQLCTARPERPMAFLREYFERLEKEEAKQIQNLQK
AGTRTDSREDEISPPPPNPVVKGRRRRGAISAEVYTEEDAASYVRKVIPKDYKTMAALAKAIEKNVLFSH
LDDNERSDIFDAMFSVSFIAGETVIQQGDEGDNFYVIDQGETDVYVNNEWATSVGEGGSFGELALIYGTP
RAATVKAKTNVKLWGIDRDSYRRILMGSTLRKRKMYEEFLSKVSILESLDKWERLTVADALEPVQFEDGQ
KIVVQGEPGDEFFIILEGSAAVLQRRSENEEFVEVGRLGPSDYFGEIALLMNRPRAATVVARGPLKCVKL
DRPRFERVLGPCSDILKRNIQQYNSFVSLSV
```

### Output
* `lba-enrichment.txt`: table with enriched terms for each prey gene