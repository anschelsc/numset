#include <stdlib.h>

#include "numset.h"

struct numset {
	unsigned int size;
	unsigned int *ordered;
	unsigned int *unordered;
};

struct numset *make_ns(unsigned int max) {
	struct numset *ret = malloc(sizeof(*ret));
	ret->size = 0;
	ret->ordered = malloc(2 * max * sizeof(*ret->ordered));
	ret->unordered = ret->ordered + max;
	return ret;
}

void free_ns(struct numset *ns) {
	free(ns->ordered);
	free(ns);
}

int get_ns(struct numset *ns, unsigned int index) {
	unsigned int u_index = ns->ordered[index];
	if (u_index < ns->size) {
		return ns->unordered[u_index] == index;
	}
	return 0;
}

void set_ns(struct numset *ns, unsigned int index) {
	ns->unordered[ns->size] = index;
	ns->ordered[index] = ns->size;
	ns->size++;
}

void unset_ns(struct numset *ns, unsigned int index) {
	unsigned int u_index = ns->ordered[index];
	if (u_index != (ns->size - 1)) {
		unsigned int other = ns->unordered[ns->size - 1];
		ns->unordered[u_index] = other;
		ns->ordered[other] = u_index;
	}
	ns->size--;
}

void clear_ns(struct numset *ns) {
	ns->size = 0;
}
