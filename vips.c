#include "vips.h"

#include <stdio.h>

/* vips */
void
vipsimage_free(VipsImage *in) {
    if (G_IS_OBJECT(in)) g_clear_object(&in);
}

/* VipsImage */
VipsImage *
vipsimage_image_new_from_file (const char *name) {
    return vips_image_new_from_file(name, NULL);
}

VipsImage *
vipsimage_image_new_from_buffer (const void *buf, size_t len, const char *option_string) {
    return vips_image_new_from_buffer(buf, len, option_string, NULL);
}

/* conversion */
int
vipsimage_replicate (VipsImage *in, VipsImage **out, int across, int down) {
    return vips_replicate(in, out, across, down, NULL);
}

int
vipsimage_gravity (VipsImage *in, VipsImage **out, VipsCompassDirection direction, int width, int height) {
    return vips_gravity(in, out, direction, width, height, NULL);
}

int
vipsimage_composite2 (VipsImage *base, VipsImage *overlay, VipsImage **out, VipsBlendMode mode, int x, int y) {
    return vips_composite2(base, overlay, out, mode, "x", x, "y", y, NULL);
}

int
vipsimage_autorot (VipsImage *in, VipsImage **out) {
    return vips_autorot(in, out, NULL);
}

int
vipsimage_embed (VipsImage *in, VipsImage **out, int x, int y, int width, int height) {
    return vips_embed(in, out, x, y, width, height, NULL);
}

int
vipsimage_extract_area (VipsImage *in, VipsImage **out, int left, int top, int width, int height) {
    return vips_extract_area(in, out, left, top, width, height, NULL);
}

int
vipsimage_crop (VipsImage *in, VipsImage **out, int left, int top, int width, int height) {
    return vips_extract_area(in, out, left, top, width, height, NULL);
}

int
vipsimage_smartcrop (VipsImage *in, VipsImage **out, int width, int height) {
    return vips_smartcrop(in, out, width, height, NULL);
}

int
vipsimage_flip (VipsImage *in, VipsImage **out, VipsDirection direction) {
    return vips_flip(in, out, direction, NULL);
}

int
vipsimage_grid (VipsImage *in, VipsImage **out, int tile_height, int across, int down) {
    return vips_grid(in, out, tile_height, across, down, NULL);
}

int
vipsimage_scale (VipsImage *in, VipsImage **out) {
    return vips_scale(in, out, NULL);
}

int
vipsimage_subsample (VipsImage *in, VipsImage **out, int xfac, int yfac) {
    return vips_subsample(in, out, xfac, yfac, NULL);
}

int
vipsimage_zoom (VipsImage *in, VipsImage **out, int xfac, int yfac) {
    return vips_zoom(in, out, xfac, yfac, NULL);
}

int
vipsimage_wrap (VipsImage *in, VipsImage **out) {
    return vips_wrap(in, out, NULL);
}

int
vipsimage_extract_band (VipsImage *in, VipsImage **out, int band) {
    return vips_extract_band(in, out, band, NULL);
}

/* VipsForeignSave */
int
vipsimage_vipsload (const char *filename, VipsImage **out) {
    return vips_vipsload(filename, out, NULL);
}

int
vipsimage_vipssave (VipsImage *in, const char *filename) {
    return vips_vipssave(in, filename, NULL);
}

int
vipsimage_openslideload (const char *filename, VipsImage **out) {
    return vips_openslideload(filename, out, NULL);
}

int
vipsimage_jpegload (const char *filename, VipsImage **out) {
    return vips_jpegload(filename, out, NULL);
}

int
vipsimage_jpegload_buffer (void *buf, size_t len, VipsImage **out) {
    return vips_jpegload_buffer(buf, len, out, NULL);
}

int
vipsimage_jpegsave (VipsImage *in, const char *filename) {
    return vips_jpegsave(in, filename, NULL);
}

int
vipsimage_jpegsave_buffer (VipsImage *in, void **buf, size_t *len) {
    return vips_jpegsave_buffer(in, buf, len, NULL);
}

int
vipsimage_jpegsave_mime (VipsImage *in) {
    return vips_jpegsave_mime(in, NULL);
}

int
vipsimage_webpload (const char *filename, VipsImage **out) {
    return vips_webpload(filename, out, NULL);
}

int
vipsimage_webpload_buffer (void *buf, size_t len, VipsImage **out) {
    return vips_webpload_buffer(buf, len, out, NULL);
}

int
vipsimage_webpsave (VipsImage *in, const char *filename) {
    return vips_webpsave(in, filename, NULL);
}

int
vipsimage_webpsave_buffer (VipsImage *in, void **buf, size_t *len) {
    return vips_webpsave_buffer(in, buf, len, NULL);
}

int
vipsimage_webpsave_mime (VipsImage *in) {
    return vips_webpsave_mime(in, NULL);
}

int
vipsimage_tiffload (const char *filename, VipsImage **out) {
    return vips_tiffload(filename, out, NULL);
}

int
vipsimage_tiffload_buffer (void *buf, size_t len, VipsImage **out) {
    return vips_tiffload_buffer(buf, len, out, NULL);
}

int
vipsimage_tiffsave (VipsImage *in, const char *filename) {
    return vips_tiffsave(in, filename, NULL);
}

int
vipsimage_tiffsave_buffer (VipsImage *in, void **buf, size_t *len) {
    return vips_tiffsave_buffer(in, buf, len, NULL);
}

int
vipsimage_openexrload (const char *filename, VipsImage **out) {
    return vips_openexrload(filename, out, NULL);
}

int
vipsimage_fitsload (const char *filename, VipsImage **out) {
    return vips_fitsload(filename, out, NULL);
}

int
vipsimage_fitssave (VipsImage *in, const char *filename) {
    return vips_fitssave(in, filename, NULL);
}

int
vipsimage_analyzeload (const char *filename, VipsImage **out) {
    return vips_analyzeload(filename, out, NULL);
}

int
vipsimage_rawload (const char *filename, VipsImage **out, int width, int height, int bands) {
    return vips_rawload(filename, out, width, height, bands, NULL);
}

int
vipsimage_rawsave (VipsImage *in, const char *filename) {
    return vips_rawsave(in, filename, NULL);
}

int
vipsimage_csvload (const char *filename, VipsImage **out) {
    return vips_csvload(filename, out, NULL);
}

int
vipsimage_csvsave (VipsImage *in, const char *filename) {
    return vips_csvsave(in, filename, NULL);
}

int
vipsimage_matrixload (const char *filename, VipsImage **out) {
    return vips_matrixload(filename, out, NULL);
}

int
vipsimage_matrixsave (VipsImage *in, const char *filename) {
    return vips_matrixsave(in, filename, NULL);
}

int
vipsimage_matrixprint (VipsImage *in) {
    return vips_matrixprint(in, NULL);
}

int
vipsimage_magickload (const char *filename, VipsImage **out) {
    return vips_magickload(filename, out, NULL);
}

int
vipsimage_magickload_buffer (void *buf, size_t len, VipsImage **out) {
    return vips_magickload_buffer(buf, len, out, NULL);
}

int
vipsimage_magicksave (VipsImage *in, const char *filename) {
    return vips_magicksave(in, filename, NULL);
}

int
vipsimage_magicksave_buffer (VipsImage *in, void **buf, size_t *len) {
    return vips_magicksave_buffer(in, buf, len, NULL);
}


int
vipsimage_pngload (const char *filename, VipsImage **out) {
    return vips_pngload(filename, out, NULL);
}

int
vipsimage_pngload_buffer (void *buf, size_t len, VipsImage **out) {
    return vips_pngload_buffer(buf, len, out, NULL);
}

int
vipsimage_pngsave (VipsImage *in, const char *filename) {
    return vips_pngsave(in, filename, NULL);
}

int
vipsimage_pngsave_buffer (VipsImage *in, void **buf, size_t *len) {
    return vips_pngsave_buffer(in, buf, len, NULL);
}

int
vipsimage_ppmload (const char *filename, VipsImage **out) {
    return vips_ppmload(filename, out, NULL);
}

int
vipsimage_ppmsave (VipsImage *in, const char *filename) {
    return vips_ppmsave(in, filename, NULL);
}

int
vipsimage_matload (const char *filename, VipsImage **out) {
    return vips_matload(filename, out, NULL);
}

int
vipsimage_radload (const char *filename, VipsImage **out) {
    return vips_radload(filename, out, NULL);
}

int
vipsimage_radsave (VipsImage *in, const char *filename) {
    return vips_radsave(in, filename, NULL);
}

int
vipsimage_radsave_buffer (VipsImage *in, void **buf, size_t *len) {
    return vips_radsave_buffer(in, buf, len, NULL);
}

int
vipsimage_pdfload (const char *filename, VipsImage **out) {
    return vips_pdfload(filename, out, NULL);
}

int
vipsimage_pdfload_buffer (void *buf, size_t len, VipsImage **out) {
    return vips_pdfload_buffer(buf, len, out, NULL);
}

int
vipsimage_svgload (const char *filename, VipsImage **out) {
    return vips_svgload(filename, out, NULL);
}

int
vipsimage_svgload_buffer (void *buf, size_t len, VipsImage **out) {
    return vips_svgload_buffer(buf, len, out, NULL);
}

int
vipsimage_gifload (const char *filename, VipsImage **out) {
    return vips_gifload(filename, out, NULL);
}

int
vipsimage_gifload_buffer (void *buf, size_t len, VipsImage **out) {
    return vips_gifload_buffer(buf, len, out, NULL);
}

int
vipsimage_heifload (const char *filename, VipsImage **out) {
    return vips_heifload(filename, out, NULL);
}

int
vipsimage_heifload_buffer (void *buf, size_t len, VipsImage **out) {
    return vips_heifload_buffer(buf, len, out, NULL);
}

int
vipsimage_heifsave (VipsImage *in, const char *filename) {
    return vips_heifsave(in, filename, NULL);
}

int
vipsimage_heifsave_buffer (VipsImage *in, void **buf, size_t *len) {
    return vips_heifsave_buffer(in, buf, len, NULL);
}

int
vipsimage_niftiload (const char *filename, VipsImage **out) {
    return vips_niftiload(filename, out, NULL);
}

int
vipsimage_niftisave (VipsImage *in, const char *filename) {
    return vips_niftisave(in, filename, NULL);
}

int
vipsimage_dzsave (VipsImage *in, const char *name) {
    return vips_dzsave(in, name, NULL);
}
/* resample */
int
vipsimage_resize (VipsImage *in, VipsImage **out, double scale) {
    return vips_resize(in, out, scale, NULL);
}

int
vipsimage_rotate (VipsImage *in, VipsImage **out, double angle) {
    return vips_rotate(in, out, angle, NULL);
}

int
vipsimage_thumbnail_image (VipsImage *in, VipsImage **out, int width) {
    return vips_thumbnail_image(in, out, width, NULL);
}

int
vipsimage_thumbnail (const char *filename, VipsImage **out, int width) {
    return vips_thumbnail(filename, out, width, NULL);
}

int
vipsimage_thumbnail_buffer (void *buf, size_t len, VipsImage **out, int width) {
    return vips_thumbnail_buffer(buf, len, out, width, NULL);
}

/* freqfilt */
int
vipsimage_fwfft (VipsImage *in, VipsImage **out) {
    return vips_fwfft(in, out, NULL);
}

int
vipsimage_invfft (VipsImage *in, VipsImage **out) {
    return vips_invfft(in, out, NULL);
}

int
vipsimage_freqmult (VipsImage *in, VipsImage *mask, VipsImage **out) {
    return vips_freqmult(in, mask, out, NULL);
}

int
vipsimage_spectrum (VipsImage *in, VipsImage **out) {
    return vips_spectrum(in, out, NULL);
}

int
vipsimage_phasecor (VipsImage *in1, VipsImage *in2, VipsImage **out) {
    return vips_phasecor(in1, in2, out, NULL);
}

/* create */
int
vipsimage_black (VipsImage **out, int width, int height) {
    return vips_black(out, width, height, NULL);
}

int
vipsimage_xyz (VipsImage **out, int width, int height) {
    return vips_xyz(out, width, height, NULL);
}

int
vipsimage_grey (VipsImage **out, int width, int height) {
    return vips_grey(out, width, height, NULL);
}

int
vipsimage_gaussmat (VipsImage **out, double sigma, double min_ampl) {
    return vips_gaussmat(out, sigma, min_ampl, NULL);
}

int
vipsimage_logmat (VipsImage **out, double sigma, double min_ampl) {
    return vips_logmat(out, sigma, min_ampl, NULL);
}

int
vipsimage_text (VipsImage **out, const char *text) {
    return vips_text(out, text, NULL);
}

int
vipsimage_gaussnoise (VipsImage **out, int width, int height) {
    return vips_gaussnoise(out, width, height, NULL);
}

int
vipsimage_eye (VipsImage **out, int width, int height) {
    return vips_eye(out, width, height, NULL);
}

int
vipsimage_sines (VipsImage **out, int width, int height) {
    return vips_sines(out, width, height, NULL);
}

int
vipsimage_zone (VipsImage **out, int width, int height) {
    return vips_zone(out, width, height, NULL);
}

int
vipsimage_identity (VipsImage **out) {
    return vips_identity(out, NULL);
}

int
vipsimage_buildlut (VipsImage *in, VipsImage **out) {
    return vips_buildlut(in, out, NULL);
}

int
vipsimage_invertlut (VipsImage *in, VipsImage **out) {
    return vips_invertlut(in, out, NULL);
}

int
vipsimage_tonelut (VipsImage **out) {
    return vips_tonelut(out, NULL);
}

int
vipsimage_mask_ideal (VipsImage **out, int width, int height, double frequency_cutoff) {
    return vips_mask_ideal(out, width, height, frequency_cutoff, NULL);
}

int
vipsimage_mask_ideal_ring (VipsImage **out, int width, int height, double frequency_cutoff, double ringwidth) {
    return vips_mask_ideal_ring(out, width, height, frequency_cutoff, ringwidth, NULL);
}

int
vipsimage_mask_ideal_band (VipsImage **out, int width, int height, double frequency_cutoff_x, double frequency_cutoff_y, double radius) {
    return vips_mask_ideal_band(out, width, height, frequency_cutoff_x, frequency_cutoff_y, radius, NULL);
}

int
vipsimage_mask_butterworth (VipsImage **out, int width, int height, double order, double frequency_cutoff, double amplitude_cutoff) {
    return vips_mask_butterworth(out, width, height, order, frequency_cutoff, amplitude_cutoff, NULL);
}

int
vipsimage_mask_butterworth_ring (VipsImage **out, int width, int height, double order, double frequency_cutoff, double amplitude_cutoff, double ringwidth) {
    return vips_mask_butterworth_ring(out, width, height, order, frequency_cutoff, amplitude_cutoff, ringwidth, NULL);
}

int
vipsimage_mask_butterworth_band (VipsImage **out,
        int width, int height,
        double order,
        double frequency_cutoff_x, double frequency_cutoff_y,
        double radius, double amplitude_cutoff) {

    return vips_mask_butterworth_band(out, width, height, order, frequency_cutoff_x, frequency_cutoff_y, radius, amplitude_cutoff, NULL);
}

int
vipsimage_mask_gaussian (VipsImage **out, int width, int height, double frequency_cutoff, double amplitude_cutoff) {
    return vips_mask_gaussian(out, width, height, frequency_cutoff, amplitude_cutoff, NULL);
}

int
vipsimage_mask_gaussian_ring (VipsImage **out, int width, int height, double frequency_cutoff, double amplitude_cutoff, double ringwidth) {
    return vips_mask_gaussian_ring(out, width, height, frequency_cutoff, amplitude_cutoff, ringwidth, NULL);
}

int
vipsimage_mask_gaussian_band (VipsImage **out, int width, int height, double frequency_cutoff_x, double frequency_cutoff_y, double radius, double amplitude_cutoff) {
    return vips_mask_gaussian_band(out, width, height, frequency_cutoff_x, frequency_cutoff_y, radius, amplitude_cutoff, NULL);
}

int
vipsimage_mask_fractal (VipsImage **out, int width, int height, double fractal_dimension) {
    return vips_mask_fractal(out, width, height, fractal_dimension, NULL);
}

int
vipsimage_fractsurf (VipsImage **out, int width, int height, double fractal_dimension) {
    return vips_fractsurf(out, width, height, fractal_dimension, NULL);
}

int
vipsimage_worley (VipsImage **out, int width, int height) {
    return vips_worley(out, width, height, NULL);
}

int
vipsimage_perlin (VipsImage **out, int width, int height) {
    return vips_perlin(out, width, height, NULL);
}


/* draw */
int
vipsimage_draw_rect (VipsImage *image, double *ink, int n, int left, int top, int width, int height) {
    return vips_draw_rect(image, ink, n, left, top, width, height, NULL);
}

int
vipsimage_draw_rect1 (VipsImage *image, double ink, int left, int top, int width, int height) {
    return vips_draw_rect1(image, ink, left, top, width, height, NULL);
}

int
vipsimage_draw_point (VipsImage *image, double *ink, int n, int x, int y) {
    return vips_draw_point(image, ink, n, x, y, NULL);
}

int
vipsimage_draw_point1 (VipsImage *image, double ink, int x, int y) {
    return vips_draw_point1(image, ink, x, y, NULL);
}

int
vipsimage_draw_image (VipsImage *image, VipsImage *sub, int x, int y) {
    return vips_draw_image(image, sub, x, y, NULL);
}

int
vipsimage_draw_mask (VipsImage *image, double *ink, int n, VipsImage *mask, int x, int y) {
    return vips_draw_mask(image, ink, n, mask, x, y, NULL);
}

int
vipsimage_draw_mask1 (VipsImage *image, double ink, VipsImage *mask, int x, int y) {
    return vips_draw_mask1(image, ink, mask, x, y, NULL);
}

int
vipsimage_draw_line (VipsImage *image, double *ink, int n, int x1, int y1, int x2, int y2) {
    return vips_draw_line(image, ink, n, x1, y1, x2, y2, NULL);
}

int
vipsimage_draw_line1 (VipsImage *image, double ink, int x1, int y1, int x2, int y2) {
    return vips_draw_line1(image, ink, x1, y1, x2, y2, NULL);
}

int
vipsimage_draw_circle (VipsImage *image, double *ink, int n, int cx, int cy, int radius) {
    return vips_draw_circle(image, ink, n, cx, cy, radius, NULL);
}

int
vipsimage_draw_circle1 (VipsImage *image, double ink, int cx, int cy, int radius) {
    return vips_draw_circle1(image, ink, cx, cy, radius, NULL);
}

int
vipsimage_draw_flood (VipsImage *image, double *ink, int n, int x, int y) {
    return vips_draw_flood(image, ink, n, x, y, NULL);
}

int
vipsimage_draw_flood1 (VipsImage *image, double ink, int x, int y) {
    return vips_draw_flood1(image, ink, x, y, NULL);
}

int
vipsimage_draw_smudge (VipsImage *image, int left, int top, int width, int height) {
    return vips_draw_smudge(image, left, top, width, height, NULL);
}


/* mosaicing */
int
vipsimage_merge (VipsImage *ref, VipsImage *sec, VipsImage **out, VipsDirection direction, int dx, int dy) {
    return vips_merge(ref, sec, out, direction, dx, dy, NULL);
}

int
vipsimage_mosaic (VipsImage *ref, VipsImage *sec, VipsImage **out, VipsDirection direction, int xref, int yref, int xsec, int ysec) {
    return vips_mosaic(ref, sec, out, direction, xref, yref, xsec, ysec, NULL);
}

int
vipsimage_mosaic1 (VipsImage *ref, VipsImage *sec, VipsImage **out, VipsDirection direction, int xr1, int yr1, int xs1, int ys1, int xr2, int yr2, int xs2, int ys2) {
    return vips_mosaic1(ref, sec, out, direction, xr1, yr1, xs1, ys1, xr2, yr2, xs2, ys2, NULL);
}

int
vipsimage_match (VipsImage *ref, VipsImage *sec, VipsImage **out, int xr1, int yr1, int xs1, int ys1, int xr2, int yr2, int xs2, int ys2) {
    return vips_match(ref, sec, out, xr1, yr1, xs1, ys1, xr2, yr2, xs2, ys2, NULL);
}

int vipsimage_globalbalance (VipsImage *in, VipsImage **out) {
    return vips_globalbalance(in, out, NULL);
}

int
vipsimage_remosaic (VipsImage *in, VipsImage **out, const char *old_str, const char *new_str) {
    return vips_remosaic(in, out, old_str, new_str, NULL);
}

int
vipsimage_copy (VipsImage *in, VipsImage **out) {
    return vips_copy(in, out, NULL);
}