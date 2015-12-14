#ifndef NUMSET
#define NUMSET

struct numset;

struct numset *make_ns(unsigned int max);

void free_ns(struct numset *);

// Not idempotent; call get_ns() first if unsure
void set_ns(struct numset *, unsigned int index);
void unset_ns(struct numset *, unsigned int index);

int get_ns(struct numset *, unsigned int index);

void clear_ns(struct numset *);

#endif
