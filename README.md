# Merge K Sorted Arrays

> Striver SDE Sheet 073 · Day 12 Heaps — a min-heap of per-array cursors; every output element costs one pop and at most one push. Go, stdlib only, tested.

## Problem

Merge k sorted arrays (total N elements) into one sorted array. Concatenate-and-sort is O(N log N) and ignores the gift that the inputs are already sorted.

## Approach

Generalize the two-list merge's "compare the heads" to k heads with a min-heap of cursors (value, which array, position). Pop the smallest overall, emit it, and push that array's next element — the heap never holds more than one cursor per array, so its size is k, not N. Each element passes through the heap once: O(N log k), strictly between the two-list merge (k=2) and heapsort (k=N). The same cursor-heap merges k sorted *lists*, files on disk, or sorted streams — it's the core of external merge sort.

**Complexity:** O(N log k) time, O(k) heap + output.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
