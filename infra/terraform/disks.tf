resource "yandex_compute_disk" "boot-disk" {
  name = "boot-disk"
  type = "network-hdd"
  zone = "ru-central1-b"
  size = "${var.boot_disk_size}"
  image_id = "fd8vnu51jg630baj500i"
}
