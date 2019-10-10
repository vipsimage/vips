#include <vips/vips.h>
/* vips */

/* VipsImage */
VipsImage *vipsimage_image_new_from_file (const char *name);
VipsImage *vipsimage_image_new_from_buffer (const void *buf, size_t len, const char *option_string);

/* conversion */
int vipsimage_replicate (VipsImage *in, VipsImage **out, int across, int down);
int vipsimage_gravity (VipsImage *in, VipsImage **out, VipsCompassDirection direction, int width, int height);
int vipsimage_composite2 (VipsImage *base, VipsImage *overlay, VipsImage **out, VipsBlendMode mode1);
int vipsimage_autorot (VipsImage *in, VipsImage **out);
int vipsimage_embed (VipsImage *in, VipsImage **out, int x, int y, int width, int height);
int vipsimage_extract_area (VipsImage *in, VipsImage **out, int left, int top, int width, int height);
int vipsimage_crop (VipsImage *in, VipsImage **out, int left, int top, int width, int height);
int vipsimage_smartcrop (VipsImage *in, VipsImage **out, int width, int height);
int vipsimage_flip (VipsImage *in, VipsImage **out, VipsDirection direction);
int vipsimage_grid (VipsImage *in, VipsImage **out, int tile_height, int across, int down);
int vipsimage_scale (VipsImage *in, VipsImage **out);
int vipsimage_subsample (VipsImage *in, VipsImage **out, int xfac, int yfac);
int vipsimage_zoom (VipsImage *in, VipsImage **out, int xfac, int yfac);
int vipsimage_wrap (VipsImage *in, VipsImage **out);
int vipsimage_extract_band (VipsImage *in, VipsImage **out, int band);
int vipsimage_copy (VipsImage *in, VipsImage **out);

/* VipsForeignSave */
int vipsimage_vipsload (const char *filename, VipsImage **out);
int vipsimage_vipssave (VipsImage *in, const char *filename);
int vipsimage_openslideload (const char *filename, VipsImage **out);
int vipsimage_jpegload (const char *filename, VipsImage **out);
int vipsimage_jpegload_buffer (void *buf, size_t len, VipsImage **out);
int vipsimage_jpegsave (VipsImage *in, const char *filename);
int vipsimage_jpegsave_buffer (VipsImage *in, void **buf, size_t *len);
int vipsimage_jpegsave_mime (VipsImage *in);
int vipsimage_webpload (const char *filename, VipsImage **out);
int vipsimage_webpload_buffer (void *buf, size_t len, VipsImage **out);
int vipsimage_webpsave (VipsImage *in, const char *filename);
int vipsimage_webpsave_buffer (VipsImage *in, void **buf, size_t *len);
int vipsimage_webpsave_mime (VipsImage *in);
int vipsimage_tiffload (const char *filename, VipsImage **out);
int vipsimage_tiffload_buffer (void *buf, size_t len, VipsImage **out);
int vipsimage_tiffsave (VipsImage *in, const char *filename);
int vipsimage_tiffsave_buffer (VipsImage *in, void **buf, size_t *len);
int vipsimage_openexrload (const char *filename, VipsImage **out);
int vipsimage_fitsload (const char *filename, VipsImage **out);
int vipsimage_fitssave (VipsImage *in, const char *filename);
int vipsimage_analyzeload (const char *filename, VipsImage **out);
int vipsimage_rawload (const char *filename, VipsImage **out, int width, int height, int bands);
int vipsimage_rawsave (VipsImage *in, const char *filename);
int vipsimage_csvload (const char *filename, VipsImage **out);
int vipsimage_csvsave (VipsImage *in, const char *filename);
int vipsimage_matrixload (const char *filename, VipsImage **out);
int vipsimage_matrixsave (VipsImage *in, const char *filename);
int vipsimage_matrixprint (VipsImage *in);
int vipsimage_magickload (const char *filename, VipsImage **out);
int vipsimage_magickload_buffer (void *buf, size_t len, VipsImage **out);
int vipsimage_magicksave (VipsImage *in, const char *filename);
int vipsimage_magicksave_buffer (VipsImage *in, void **buf, size_t *len);
int vipsimage_pngload (const char *filename, VipsImage **out);
int vipsimage_pngload_buffer (void *buf, size_t len, VipsImage **out);
int vipsimage_pngsave (VipsImage *in, const char *filename);
int vipsimage_pngsave_buffer (VipsImage *in, void **buf, size_t *len);
int vipsimage_ppmload (const char *filename, VipsImage **out);
int vipsimage_ppmsave (VipsImage *in, const char *filename);
int vipsimage_matload (const char *filename, VipsImage **out);
int vipsimage_radload (const char *filename, VipsImage **out);
int vipsimage_radsave (VipsImage *in, const char *filename);
int vipsimage_radsave_buffer (VipsImage *in, void **buf, size_t *len);
int vipsimage_pdfload (const char *filename, VipsImage **out);
int vipsimage_pdfload_buffer (void *buf, size_t len, VipsImage **out);
int vipsimage_svgload (const char *filename, VipsImage **out);
int vipsimage_svgload_buffer (void *buf, size_t len, VipsImage **out);
int vipsimage_gifload (const char *filename, VipsImage **out);
int vipsimage_gifload_buffer (void *buf, size_t len, VipsImage **out);
int vipsimage_heifload (const char *filename, VipsImage **out);
int vipsimage_heifload_buffer (void *buf, size_t len, VipsImage **out);
int vipsimage_heifsave (VipsImage *in, const char *filename);
int vipsimage_heifsave_buffer (VipsImage *in, void **buf, size_t *len);
int vipsimage_niftiload (const char *filename, VipsImage **out);
int vipsimage_niftisave (VipsImage *in, const char *filename);
int vipsimage_dzsave (VipsImage *in, const char *name);
/* resample */
int vipsimage_resize (VipsImage *in, VipsImage **out, double scale);
int vipsimage_rotate (VipsImage *in, VipsImage **out, double angle);
int vipsimage_thumbnail_image (VipsImage *in, VipsImage **out, int width);
int vipsimage_thumbnail (const char *filename, VipsImage **out, int width);
int vipsimage_thumbnail_buffer (void *buf, size_t len, VipsImage **out, int width);

/* freqfilt */
int vipsimage_fwfft (VipsImage *in, VipsImage **out);
int vipsimage_invfft (VipsImage *in, VipsImage **out);
int vipsimage_freqmult (VipsImage *in, VipsImage *mask, VipsImage **out);
int vipsimage_spectrum (VipsImage *in, VipsImage **out);
int vipsimage_phasecor (VipsImage *in1, VipsImage *in2, VipsImage **out);

/* create */
int vipsimage_black (VipsImage **out, int width, int height);
int vipsimage_xyz (VipsImage **out, int width, int height);
int vipsimage_grey (VipsImage **out, int width, int height);
int vipsimage_gaussmat (VipsImage **out, double sigma, double min_ampl);
int vipsimage_logmat (VipsImage **out, double sigma, double min_ampl);
int vipsimage_text (VipsImage **out, const char *text);
int vipsimage_gaussnoise (VipsImage **out, int width, int height);
int vipsimage_eye (VipsImage **out, int width, int height);
int vipsimage_sines (VipsImage **out, int width, int height);
int vipsimage_zone (VipsImage **out, int width, int height);
int vipsimage_identity (VipsImage **out);
int vipsimage_buildlut (VipsImage *in, VipsImage **out);
int vipsimage_invertlut (VipsImage *in, VipsImage **out);
int vipsimage_tonelut (VipsImage **out);
int vipsimage_mask_ideal (VipsImage **out, int width, int height, double frequency_cutoff);
int vipsimage_mask_ideal_ring (VipsImage **out, int width, int height, double frequency_cutoff, double ringwidth);
int vipsimage_mask_ideal_band (VipsImage **out, int width, int height, double frequency_cutoff_x, double frequency_cutoff_y, double radius);
int vipsimage_mask_butterworth (VipsImage **out, int width, int height, double order, double frequency_cutoff, double amplitude_cutoff);
int vipsimage_mask_butterworth_ring (VipsImage **out, int width, int height, double order, double frequency_cutoff, double amplitude_cutoff, double ringwidth);
int vipsimage_mask_butterworth_band (VipsImage **out, int width, int height, double order, double frequency_cutoff_x, double frequency_cutoff_y, double radius, double amplitude_cutoff);
int vipsimage_mask_gaussian (VipsImage **out, int width, int height, double frequency_cutoff, double amplitude_cutoff);
int vipsimage_mask_gaussian_ring (VipsImage **out, int width, int height, double frequency_cutoff, double amplitude_cutoff, double ringwidth);
int vipsimage_mask_gaussian_band (VipsImage **out, int width, int height, double frequency_cutoff_x, double frequency_cutoff_y, double radius, double amplitude_cutoff);
int vipsimage_mask_fractal (VipsImage **out, int width, int height, double fractal_dimension);
int vipsimage_fractsurf (VipsImage **out, int width, int height, double fractal_dimension);
int vipsimage_worley (VipsImage **out, int width, int height);
int vipsimage_perlin (VipsImage **out, int width, int height);

/* draw */
int vipsimage_draw_rect (VipsImage *image, double *ink, int n, int left, int top, int width, int height);
int vipsimage_draw_rect1 (VipsImage *image, double ink, int left, int top, int width, int height);
int vipsimage_draw_point (VipsImage *image, double *ink, int n, int x, int y);
int vipsimage_draw_point1 (VipsImage *image, double ink, int x, int y);
int vipsimage_draw_image (VipsImage *image, VipsImage *sub, int x, int y);
int vipsimage_draw_mask (VipsImage *image, double *ink, int n, VipsImage *mask, int x, int y);
int vipsimage_draw_mask1 (VipsImage *image, double ink, VipsImage *mask, int x, int y);
int vipsimage_draw_line (VipsImage *image, double *ink, int n, int x1, int y1, int x2, int y2);
int vipsimage_draw_line1 (VipsImage *image, double ink, int x1, int y1, int x2, int y2);
int vipsimage_draw_circle (VipsImage *image, double *ink, int n, int cx, int cy, int radius);
int vipsimage_draw_circle1 (VipsImage *image, double ink, int cx, int cy, int radius);
int vipsimage_draw_flood (VipsImage *image, double *ink, int n, int x, int y);
int vipsimage_draw_flood1 (VipsImage *image, double ink, int x, int y);
int vipsimage_draw_smudge (VipsImage *image, int left, int top, int width, int height);

/* mosaicing */
int vipsimage_merge (VipsImage *ref, VipsImage *sec, VipsImage **out, VipsDirection direction, int dx, int dy);
int vipsimage_mosaic (VipsImage *ref, VipsImage *sec, VipsImage **out, VipsDirection direction, int xref, int yref, int xsec, int ysec);
int vipsimage_mosaic1 (VipsImage *ref, VipsImage *sec, VipsImage **out, VipsDirection direction, int xr1, int yr1, int xs1, int ys1, int xr2, int yr2, int xs2, int ys2);
int vipsimage_match (VipsImage *ref, VipsImage *sec, VipsImage **out, int xr1, int yr1, int xs1, int ys1, int xr2, int yr2, int xs2, int ys2);
int vipsimage_globalbalance (VipsImage *in, VipsImage **out);
int vipsimage_remosaic (VipsImage *in, VipsImage **out, const char *old_str, const char *new_str);
