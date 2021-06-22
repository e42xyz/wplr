<?php

/**
 * My customizations for adding livereload and hooking into various WordPress events 
 * that users may want to trigger a livereload (additional hooks to be added).
 * 
 * @link https://hakk.dev/wplr/
 * @since 5.7.2
 * @package wplr-plugin
 * 
 * * @wordpress-plugin
 * Plugin Name:       WordPress LiveReload Plugin
 * Plugin URI:        https://hakk.dev/wplr/
 * Description:       This plugin can help with site development by automatically reloading the page when a file changes on disk.
 * Version:           1.0.0
 * Author:            bmcculley
 * Author URI:        https://hakk.dev
 * License:           GPL-2.0+
 * License URI:       http://www.gnu.org/licenses/gpl-2.0.txt
 * Text Domain:       wplr-plugin
 */

// If this file is called directly, abort.
if ( ! defined( 'WPINC' ) ) {
	die;
}

/**
 * Current plugin version.
 */
define( 'PLUGIN_NAME_VERSION', '1.0.0' );

/**
 * if site_url contains a port number; remove it
 */
function remove_port($site_url) {
    $url_parts = parse_url($site_url);
    return $url_parts['scheme'].'://'.$url_parts['host'];
}

/**
 * add livereload script to wp head
 */
function wpcld_livereload_js() {
    $site_url = get_site_url();
    // check if the site url contains a :, if it does assume that the url contains a port
    if (strpos($site_url, ':') !== false){
        $site_url = remove_port($site_url);
    }
    
    echo '<script type="text/javascript" src="'.$site_url.':35729/livereload.js"></script>';
}

// Add hook for front-end <head></head>
add_action( 'wp_head', 'wpcld_livereload_js' );

?>