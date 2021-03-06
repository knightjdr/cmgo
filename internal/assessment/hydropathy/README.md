# Module name: `assessment-hydropathy`

> Measures the average hydropathy of proteins in a SAINT file and a BioPlex dataset

| parameters | definition | default |
|------------|------------|---------|
| bioplexFile | BioPlex dataset | |
| database | sequence database in fasta format | |
| FDR | FDR threshold for significant interactors | 0.01 |
| saintFile | SAINT file | |

## Example file formats

### bioplexFile
```
GeneA	GeneB	UniprotA	UniprotB	SymbolA	SymbolB	p(Wrong)	p(No Interaction)	p(Interaction)
100	728378	P00813	A5A3E0	ADA	POTEF	2.38085788908859e-9	0.000331855941652957	0.999668141677489
100	345651	P00813	Q562R1	ADA	ACTBL2	9.78643725788521e-18	0.211914436568748	0.788085563431252
222389	708	Q8N7W2	Q07021	BEND7	C1QBP	2.96221526059856e-17	0.00564451180569955	0.994355488194301
222389	4038	Q8N7W2	O75096	BEND7	LRP4	3.30299445393738e-10	0.000280259555661228	0.999719740114039
645121	3312	Q6ZMN8	P11142	CCNI2	HSPA8	2.06028533960837e-16	0.0362347656743182	0.963765234325682
645121	55132	Q6ZMN8	Q659C4-2	CCNI2	LARP1B	1.54566627347441e-12	0.0315484115050354	0.968451588493419
645121	1020	Q6ZMN8	Q00535	CCNI2	CDK5	7.11903595597094e-10	0.0000638718598341267	0.999936127428262
945	201266	P20138	Q8N1S5	CD33	SLC39A11	3.13297318944057e-11	0.0864322481626163	0.913567751806054
```

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
* stdout