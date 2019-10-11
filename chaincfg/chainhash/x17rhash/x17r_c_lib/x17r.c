/**
 * X17R algorithm (X17 with Randomized chain order)
 *
 * tpruvot 2018 - GPL code
 */

#include <stdio.h>
#include <memory.h>

#include <stdint.h>

#include "sph_blake.h"
#include "sph_bmw.h"
#include "sph_groestl.h"
#include "sph_skein.h"
#include "sph_jh.h"
#include "sph_keccak.h"

#include "sph_luffa.h"
#include "sph_cubehash.h"
#include "sph_shavite.h"
#include "sph_simd.h"
#include "sph_echo.h"

#include "sph_hamsi.h"
#include "sph_fugue.h"
#include "sph_shabal.h"
#include "sph_whirlpool.h"
#include "sph_sha2.h"
#include "compat.h"

#define MAX_GPUS 16

enum Algo {
	BLAKE = 0,
	BMW,
	GROESTL,
	JH,
	KECCAK,
	SKEIN,
	LUFFA,
	CUBEHASH,
	SHAVITE,
	SIMD,
	ECHO,
	HAMSI,
	FUGUE,
	SHABAL,
	WHIRLPOOL,
	SHA512,
	HASH_FUNC_COUNT
};

static const char* algo_strings[] = {
	"blake",
	"bmw512",
	"groestl",
	"jh512",
	"keccak",
	"skein",
	"luffa",
	"cube",
	"shavite",
	"simd",
	"echo",
	"hamsi",
	"fugue",
	"shabal",
	"whirlpool",
	"sha512",
	NULL
};

static __thread char hashOrder[HASH_FUNC_COUNT + 1] = { 0 };

void x17r_hash_impl(void *output, const void *input);

static void be32enc(void *pp, uint32_t x)
{
	uint8_t *p = (uint8_t *)pp;
	p[3] = x & 0xff;
	p[2] = (x >> 8) & 0xff;
	p[1] = (x >> 16) & 0xff;
	p[0] = (x >> 24) & 0xff;
}

static void getAlgoString(const uint32_t* prevblock, char *output)
{
	char *sptr = output;
	uint8_t* data = (uint8_t*)prevblock;

	for (uint8_t j = 0; j < HASH_FUNC_COUNT; j++) {
		uint8_t b = (15 - j) >> 1; // 16 ascii hex chars, reversed
		uint8_t algoDigit = (j & 1) ? data[b] & 0xF : data[b] >> 4;
		if (algoDigit >= 10)
			sprintf(sptr, "%c", 'A' + (algoDigit - 10));
		else
			sprintf(sptr, "%u", (uint32_t) algoDigit);
		sptr++;
	}
	*sptr = '\0';
}

#define test 70
void x17r_hash(void* output, void* input, const int in_len)
{

	//unsigned char *pdata = input;
	//unsigned char *po = output;
	//*(po + 0) = *(pdata + test - 0);
	//
	//*(po+1) = *(pdata + test - 1);
	//*(po + 2) = *(pdata + test - 2);
	//*(po + 3) = *(pdata + test - 3);
	//*(po + 4) = *(pdata + test - 4);
	//*(po + 5) = *(pdata + test - 5);
	//*(po + 6) = *(pdata + test - 6);
	//*(po + 7) = *(pdata + test - 7);
	//*(po + 8) = *(pdata + test - 8);
	//*(po + 9) = *(pdata + test - 9);

	//return;


	uint32_t *pdata = input;

	uint32_t _ALIGN(64) endiandata[20];
	for (int k = 0; k < 20; k++)
		be32enc(&endiandata[k], pdata[k]);

	uint32_t _ALIGN(64) vhash[8];
	x17r_hash_impl(vhash, endiandata);
	if (output)
	{
		memcpy(output, &vhash, 32);
	}

}

//uint32_t s_ntime = UINT32_MAX;
void x17r_hash_impl(void *output, const void *input)
{
	unsigned char _ALIGN(64) hash[128];

	sph_blake512_context ctx_blake;
	sph_bmw512_context ctx_bmw;
	sph_groestl512_context ctx_groestl;
	sph_jh512_context ctx_jh;
	sph_keccak512_context ctx_keccak;
	sph_skein512_context ctx_skein;
	sph_luffa512_context ctx_luffa;
	sph_cubehash512_context ctx_cubehash;
	sph_shavite512_context ctx_shavite;
	sph_simd512_context ctx_simd;
	sph_echo512_context ctx_echo;
	sph_hamsi512_context ctx_hamsi;
	sph_fugue512_context ctx_fugue;
	sph_shabal512_context ctx_shabal;
	sph_whirlpool_context ctx_whirlpool;
	sph_sha512_context ctx_sha512;

	void *in = (void*) input;
	int size = 80;

	uint32_t *in32 = (uint32_t*) input;
	getAlgoString(&in32[1], hashOrder);

	for (int i = 0; i < 16; i++)
	{
		const char elem = hashOrder[i];
		const uint8_t algo = elem >= 'A' ? elem - 'A' + 10 : elem - '0';

		switch (algo) {
		case BLAKE:
			sph_blake512_init(&ctx_blake);
			sph_blake512(&ctx_blake, in, size);
			sph_blake512_close(&ctx_blake, hash);
			break;
		case BMW:
			sph_bmw512_init(&ctx_bmw);
			sph_bmw512(&ctx_bmw, in, size);
			sph_bmw512_close(&ctx_bmw, hash);
			break;
		case GROESTL:
			sph_groestl512_init(&ctx_groestl);
			sph_groestl512(&ctx_groestl, in, size);
			sph_groestl512_close(&ctx_groestl, hash);
			break;
		case SKEIN:
			sph_skein512_init(&ctx_skein);
			sph_skein512(&ctx_skein, in, size);
			sph_skein512_close(&ctx_skein, hash);
			break;
		case JH:
			sph_jh512_init(&ctx_jh);
			sph_jh512(&ctx_jh, in, size);
			sph_jh512_close(&ctx_jh, hash);
			break;
		case KECCAK:
			sph_keccak512_init(&ctx_keccak);
			sph_keccak512(&ctx_keccak, in, size);
			sph_keccak512_close(&ctx_keccak, hash);
			break;
		case LUFFA:
			sph_luffa512_init(&ctx_luffa);
			sph_luffa512(&ctx_luffa, in, size);
			sph_luffa512_close(&ctx_luffa, hash);
			break;
		case CUBEHASH:
			sph_cubehash512_init(&ctx_cubehash);
			sph_cubehash512(&ctx_cubehash, in, size);
			sph_cubehash512_close(&ctx_cubehash, hash);
			break;
		case SHAVITE:
			sph_shavite512_init(&ctx_shavite);
			sph_shavite512(&ctx_shavite, in, size);
			sph_shavite512_close(&ctx_shavite, hash);
			break;
		case SIMD:
			sph_simd512_init(&ctx_simd);
			sph_simd512(&ctx_simd, in, size);
			sph_simd512_close(&ctx_simd, hash);
			break;
		case ECHO:
			sph_echo512_init(&ctx_echo);
			sph_echo512(&ctx_echo, in, size);
			sph_echo512_close(&ctx_echo, hash);
			break;
		case HAMSI:
			sph_hamsi512_init(&ctx_hamsi);
			sph_hamsi512(&ctx_hamsi, in, size);
			sph_hamsi512_close(&ctx_hamsi, hash);
			break;
		case FUGUE:
			sph_fugue512_init(&ctx_fugue);
			sph_fugue512(&ctx_fugue, in, size);
			sph_fugue512_close(&ctx_fugue, hash);
			break;
		case SHABAL:
			sph_shabal512_init(&ctx_shabal);
			sph_shabal512(&ctx_shabal, in, size);
			sph_shabal512_close(&ctx_shabal, hash);
			break;
		case WHIRLPOOL:
			sph_whirlpool_init(&ctx_whirlpool);
			sph_whirlpool(&ctx_whirlpool, in, size);
			sph_whirlpool_close(&ctx_whirlpool, hash);
			break;
		case SHA512:
			sph_sha512_init(&ctx_sha512);
			sph_sha512(&ctx_sha512,(const void*) in, size);
			sph_sha512_close(&ctx_sha512,(void*) hash);
			break;
		}
		in = (void*) hash;
		size = 64;
	}
	memcpy(output, hash, 32);
}

