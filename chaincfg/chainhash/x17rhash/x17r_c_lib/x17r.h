#ifndef __X17R_H__
#define __X17R_H__
#ifdef __cplusplus
extern "C" {
#endif

#include <stdint.h>

#include <stddef.h>

void x17r_hash(void* output, void* input, const int in_len);
#ifdef __cplusplus
}
#endif

#endif /* __X17R_H__ */
