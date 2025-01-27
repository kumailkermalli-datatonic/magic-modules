package alloydb_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-google/google/acctest"
)

func TestAccAlloydbUser_updateRoles_BuiltIn(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAlloydbUserDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAlloydbUser_alloydbUserBuiltinExample(context),
			},
			{
				ResourceName:            "google_alloydb_user.user1",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			{
				Config: testAccAlloydbUser_updateRoles_BuiltIn(context),
			},
			{
				ResourceName:            "google_alloydb_user.user1",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
		},
	})
}

func testAccAlloydbUser_updateRoles_BuiltIn(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_alloydb_instance" "default" {
  cluster       = google_alloydb_cluster.default.name
  instance_id   = "tf-test-alloydb-instance%{random_suffix}"
  instance_type = "PRIMARY"

  depends_on = [google_service_networking_connection.vpc_connection]
}

resource "google_alloydb_cluster" "default" {
  cluster_id = "tf-test-alloydb-cluster%{random_suffix}"
  location   = "us-central1"
  network    = google_compute_network.default.id

  initial_user {
    password = "tf_test_cluster_secret%{random_suffix}"
  }
}

data "google_project" "project" {}

resource "google_compute_network" "default" {
  name = "tf-test-alloydb-network%{random_suffix}"
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          = "tf-test-alloydb-cluster%{random_suffix}"
  address_type  = "INTERNAL"
  purpose       = "VPC_PEERING"
  prefix_length = 16
  network       = google_compute_network.default.id
}

resource "google_service_networking_connection" "vpc_connection" {
  network                 = google_compute_network.default.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
}

resource "google_alloydb_user" "user1" {
  cluster = google_alloydb_cluster.default.name
  user_id = "user1%{random_suffix}"
  user_type = "ALLOYDB_BUILT_IN"

  password = "tf_test_user_secret%{random_suffix}"
  database_roles = []
  depends_on = [google_alloydb_instance.default]
}`, context)
}

func TestAccAlloydbUser_updatePassword_BuiltIn(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAlloydbUserDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAlloydbUser_alloydbUserBuiltinExample(context),
			},
			{
				ResourceName:            "google_alloydb_user.user1",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
			{
				Config: testAccAlloydbUser_updatePass_BuiltIn(context),
			},
			{
				ResourceName:            "google_alloydb_user.user1",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"password"},
			},
		},
	})
}

func testAccAlloydbUser_updatePass_BuiltIn(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_alloydb_instance" "default" {
  cluster       = google_alloydb_cluster.default.name
  instance_id   = "tf-test-alloydb-instance%{random_suffix}"
  instance_type = "PRIMARY"

  depends_on = [google_service_networking_connection.vpc_connection]
}

resource "google_alloydb_cluster" "default" {
  cluster_id = "tf-test-alloydb-cluster%{random_suffix}"
  location   = "us-central1"
  network    = google_compute_network.default.id

  initial_user {
    password = "tf_test_cluster_secret%{random_suffix}"
  }
}

data "google_project" "project" {}

resource "google_compute_network" "default" {
  name = "tf-test-alloydb-network%{random_suffix}"
}

resource "google_compute_global_address" "private_ip_alloc" {
  name          = "tf-test-alloydb-cluster%{random_suffix}"
  address_type  = "INTERNAL"
  purpose       = "VPC_PEERING"
  prefix_length = 16
  network       = google_compute_network.default.id
}

resource "google_service_networking_connection" "vpc_connection" {
  network                 = google_compute_network.default.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
}

resource "google_alloydb_user" "user1" {
  cluster = google_alloydb_cluster.default.name
  user_id = "user1%{random_suffix}"
  user_type = "ALLOYDB_BUILT_IN"

  password = "tf_test_user_secret%{random_suffix}-foo"
  database_roles = ["alloydbsuperuser"]
  depends_on = [google_alloydb_instance.default]
}`, context)
}

func TestAccAlloydbUser_updateRoles_IAM(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckAlloydbUserDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccAlloydbUser_alloydbUserIamExample(context),
			},
			{
				ResourceName:            "google_alloydb_user.user2",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
			{
				Config: testAccAlloydbUser_updateRoles_Iam(context),
			},
			{
				ResourceName:            "google_alloydb_user.user2",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
			},
		},
	})
}

func testAccAlloydbUser_updateRoles_Iam(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_alloydb_instance" "default" {
  cluster       = google_alloydb_cluster.default.name
  instance_id   = "tf-test-alloydb-instance%{random_suffix}"
  instance_type = "PRIMARY"
  depends_on = [google_service_networking_connection.vpc_connection]
}
resource "google_alloydb_cluster" "default" {
  cluster_id = "tf-test-alloydb-cluster%{random_suffix}"
  location   = "us-central1"
  network    = google_compute_network.default.id
  initial_user {
    password = "tf_test_cluster_secret%{random_suffix}"
  }
}
data "google_project" "project" {}
resource "google_compute_network" "default" {
  name = "tf-test-alloydb-network%{random_suffix}"
}
resource "google_compute_global_address" "private_ip_alloc" {
  name          = "tf-test-alloydb-cluster%{random_suffix}"
  address_type  = "INTERNAL"
  purpose       = "VPC_PEERING"
  prefix_length = 16
  network       = google_compute_network.default.id
}
resource "google_service_networking_connection" "vpc_connection" {
  network                 = google_compute_network.default.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.private_ip_alloc.name]
}
resource "google_alloydb_user" "user2" {
  cluster = google_alloydb_cluster.default.name
  user_id = "user2@foo.com%{random_suffix}"
  user_type = "ALLOYDB_IAM_USER"
  database_roles = ["alloydbiamuser", "alloydbsuperuser"]
  depends_on = [google_alloydb_instance.default]
}`, context)
}
