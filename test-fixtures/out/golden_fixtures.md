# (1/64) noop

```
<nil>
```
# (2/64) base

```
<nil>
```
# (3/64) changed

```
1 errors:
changed file: test-fixtures/tmp/changed/changed.txt

run `GOLDY=update go test` to automatically update all files above
```
# (4/64) base\_changed

```
1 errors:
changed file: test-fixtures/tmp/base_changed/changed.txt

run `GOLDY=update go test` to automatically update all files above
```
# (5/64) ignore

```
<nil>
```
# (6/64) base\_ignore

```
<nil>
```
# (7/64) changed\_ignore

```
1 errors:
changed file: test-fixtures/tmp/changed_ignore/changed.txt

run `GOLDY=update go test` to automatically update all files above
```
# (8/64) base\_changed\_ignore

```
1 errors:
changed file: test-fixtures/tmp/base_changed_ignore/changed.txt

run `GOLDY=update go test` to automatically update all files above
```
# (9/64) missing

```
1 errors:
missing file: test-fixtures/tmp/missing/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (10/64) base\_missing

```
1 errors:
missing file: test-fixtures/tmp/base_missing/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (11/64) changed\_missing

```
2 errors:
changed file: test-fixtures/tmp/changed_missing/changed.txt
missing file: test-fixtures/tmp/changed_missing/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (12/64) base\_changed\_missing

```
2 errors:
changed file: test-fixtures/tmp/base_changed_missing/changed.txt
missing file: test-fixtures/tmp/base_changed_missing/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (13/64) ignore\_missing

```
1 errors:
missing file: test-fixtures/tmp/ignore_missing/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (14/64) base\_ignore\_missing

```
1 errors:
missing file: test-fixtures/tmp/base_ignore_missing/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (15/64) changed\_ignore\_missing

```
2 errors:
changed file: test-fixtures/tmp/changed_ignore_missing/changed.txt
missing file: test-fixtures/tmp/changed_ignore_missing/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (16/64) base\_changed\_ignore\_missing

```
2 errors:
changed file: test-fixtures/tmp/base_changed_ignore_missing/changed.txt
missing file: test-fixtures/tmp/base_changed_ignore_missing/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (17/64) unexpected

```
1 errors:
unexpected file: test-fixtures/tmp/unexpected/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (18/64) base\_unexpected

```
1 errors:
unexpected file: test-fixtures/tmp/base_unexpected/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (19/64) changed\_unexpected

```
2 errors:
changed file: test-fixtures/tmp/changed_unexpected/changed.txt
unexpected file: test-fixtures/tmp/changed_unexpected/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (20/64) base\_changed\_unexpected

```
2 errors:
changed file: test-fixtures/tmp/base_changed_unexpected/changed.txt
unexpected file: test-fixtures/tmp/base_changed_unexpected/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (21/64) ignore\_unexpected

```
<nil>
```
# (22/64) base\_ignore\_unexpected

```
<nil>
```
# (23/64) changed\_ignore\_unexpected

```
1 errors:
changed file: test-fixtures/tmp/changed_ignore_unexpected/changed.txt

run `GOLDY=update go test` to automatically update all files above
```
# (24/64) base\_changed\_ignore\_unexpected

```
1 errors:
changed file: test-fixtures/tmp/base_changed_ignore_unexpected/changed.txt

run `GOLDY=update go test` to automatically update all files above
```
# (25/64) missing\_unexpected

```
2 errors:
missing file: test-fixtures/tmp/missing_unexpected/missing.txt
unexpected file: test-fixtures/tmp/missing_unexpected/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (26/64) base\_missing\_unexpected

```
2 errors:
missing file: test-fixtures/tmp/base_missing_unexpected/missing.txt
unexpected file: test-fixtures/tmp/base_missing_unexpected/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (27/64) changed\_missing\_unexpected

```
3 errors:
changed file: test-fixtures/tmp/changed_missing_unexpected/changed.txt
missing file: test-fixtures/tmp/changed_missing_unexpected/missing.txt
unexpected file: test-fixtures/tmp/changed_missing_unexpected/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (28/64) base\_changed\_missing\_unexpected

```
3 errors:
changed file: test-fixtures/tmp/base_changed_missing_unexpected/changed.txt
missing file: test-fixtures/tmp/base_changed_missing_unexpected/missing.txt
unexpected file: test-fixtures/tmp/base_changed_missing_unexpected/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (29/64) ignore\_missing\_unexpected

```
1 errors:
missing file: test-fixtures/tmp/ignore_missing_unexpected/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (30/64) base\_ignore\_missing\_unexpected

```
1 errors:
missing file: test-fixtures/tmp/base_ignore_missing_unexpected/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (31/64) changed\_ignore\_missing\_unexpected

```
2 errors:
changed file: test-fixtures/tmp/changed_ignore_missing_unexpected/changed.txt
missing file: test-fixtures/tmp/changed_ignore_missing_unexpected/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (32/64) base\_changed\_ignore\_missing\_unexpected

```
2 errors:
changed file: test-fixtures/tmp/base_changed_ignore_missing_unexpected/changed.txt
missing file: test-fixtures/tmp/base_changed_ignore_missing_unexpected/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (33/64) diff

```
<nil>
```
# (34/64) base\_diff

```
<nil>
```
# (35/64) changed\_diff

```
1 errors:
changed file: test-fixtures/tmp/changed_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt

run `GOLDY=update go test` to automatically update all files above
```
# (36/64) base\_changed\_diff

```
1 errors:
changed file: test-fixtures/tmp/base_changed_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt

run `GOLDY=update go test` to automatically update all files above
```
# (37/64) ignore\_diff

```
<nil>
```
# (38/64) base\_ignore\_diff

```
<nil>
```
# (39/64) changed\_ignore\_diff

```
1 errors:
changed file: test-fixtures/tmp/changed_ignore_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt

run `GOLDY=update go test` to automatically update all files above
```
# (40/64) base\_changed\_ignore\_diff

```
1 errors:
changed file: test-fixtures/tmp/base_changed_ignore_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt

run `GOLDY=update go test` to automatically update all files above
```
# (41/64) missing\_diff

```
1 errors:
missing file: test-fixtures/tmp/missing_diff/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (42/64) base\_missing\_diff

```
1 errors:
missing file: test-fixtures/tmp/base_missing_diff/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (43/64) changed\_missing\_diff

```
2 errors:
changed file: test-fixtures/tmp/changed_missing_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt
missing file: test-fixtures/tmp/changed_missing_diff/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (44/64) base\_changed\_missing\_diff

```
2 errors:
changed file: test-fixtures/tmp/base_changed_missing_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt
missing file: test-fixtures/tmp/base_changed_missing_diff/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (45/64) ignore\_missing\_diff

```
1 errors:
missing file: test-fixtures/tmp/ignore_missing_diff/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (46/64) base\_ignore\_missing\_diff

```
1 errors:
missing file: test-fixtures/tmp/base_ignore_missing_diff/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (47/64) changed\_ignore\_missing\_diff

```
2 errors:
changed file: test-fixtures/tmp/changed_ignore_missing_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt
missing file: test-fixtures/tmp/changed_ignore_missing_diff/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (48/64) base\_changed\_ignore\_missing\_diff

```
2 errors:
changed file: test-fixtures/tmp/base_changed_ignore_missing_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt
missing file: test-fixtures/tmp/base_changed_ignore_missing_diff/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (49/64) unexpected\_diff

```
1 errors:
unexpected file: test-fixtures/tmp/unexpected_diff/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (50/64) base\_unexpected\_diff

```
1 errors:
unexpected file: test-fixtures/tmp/base_unexpected_diff/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (51/64) changed\_unexpected\_diff

```
2 errors:
changed file: test-fixtures/tmp/changed_unexpected_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt
unexpected file: test-fixtures/tmp/changed_unexpected_diff/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (52/64) base\_changed\_unexpected\_diff

```
2 errors:
changed file: test-fixtures/tmp/base_changed_unexpected_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt
unexpected file: test-fixtures/tmp/base_changed_unexpected_diff/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (53/64) ignore\_unexpected\_diff

```
<nil>
```
# (54/64) base\_ignore\_unexpected\_diff

```
<nil>
```
# (55/64) changed\_ignore\_unexpected\_diff

```
1 errors:
changed file: test-fixtures/tmp/changed_ignore_unexpected_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt

run `GOLDY=update go test` to automatically update all files above
```
# (56/64) base\_changed\_ignore\_unexpected\_diff

```
1 errors:
changed file: test-fixtures/tmp/base_changed_ignore_unexpected_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt

run `GOLDY=update go test` to automatically update all files above
```
# (57/64) missing\_unexpected\_diff

```
2 errors:
missing file: test-fixtures/tmp/missing_unexpected_diff/missing.txt
unexpected file: test-fixtures/tmp/missing_unexpected_diff/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (58/64) base\_missing\_unexpected\_diff

```
2 errors:
missing file: test-fixtures/tmp/base_missing_unexpected_diff/missing.txt
unexpected file: test-fixtures/tmp/base_missing_unexpected_diff/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (59/64) changed\_missing\_unexpected\_diff

```
3 errors:
changed file: test-fixtures/tmp/changed_missing_unexpected_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt
missing file: test-fixtures/tmp/changed_missing_unexpected_diff/missing.txt
unexpected file: test-fixtures/tmp/changed_missing_unexpected_diff/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (60/64) base\_changed\_missing\_unexpected\_diff

```
3 errors:
changed file: test-fixtures/tmp/base_changed_missing_unexpected_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt
missing file: test-fixtures/tmp/base_changed_missing_unexpected_diff/missing.txt
unexpected file: test-fixtures/tmp/base_changed_missing_unexpected_diff/unexpected.txt

run `GOLDY=update go test` to automatically update all files above
```
# (61/64) ignore\_missing\_unexpected\_diff

```
1 errors:
missing file: test-fixtures/tmp/ignore_missing_unexpected_diff/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (62/64) base\_ignore\_missing\_unexpected\_diff

```
1 errors:
missing file: test-fixtures/tmp/base_ignore_missing_unexpected_diff/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (63/64) changed\_ignore\_missing\_unexpected\_diff

```
2 errors:
changed file: test-fixtures/tmp/changed_ignore_missing_unexpected_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt
missing file: test-fixtures/tmp/changed_ignore_missing_unexpected_diff/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
# (64/64) base\_changed\_ignore\_missing\_unexpected\_diff

```
2 errors:
changed file: test-fixtures/tmp/base_changed_ignore_missing_unexpected_diff/changed.txt
  @@ -1 +1 @@
  -data for: changed.txt
  +changed data for: changed.txt
missing file: test-fixtures/tmp/base_changed_ignore_missing_unexpected_diff/missing.txt

run `GOLDY=update go test` to automatically update all files above
```
