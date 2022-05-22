Fluroclock-clockupdater
====================

Checks for the existence of `/etc/fclock-clock-enable`, then makes a HTTP POST to the Update URL from the envvar `UPDATE_URL`, then exits.

Run in a systemd timer job every minute to update the clock display with the correct time.

