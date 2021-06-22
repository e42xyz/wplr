# WordPress Live Reload

My customizations for adding livereload and hooking into various WordPress events that users may want to trigger a livereload (additional hooks to be added). 

## Usage

 - Install WordPress.
 - Download repo.
 - Compile main.go
 - Copy wplr-plugin to the /wp-content/plugins directory.
 - Activate the plugin.
 - Run `./main -p /path/to/wordpress/install`

Whenever a file is saved the site will automatically reload.
