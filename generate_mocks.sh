create_mock() {
		mockgen -source=$1.go -destination=gotable_mock/mock_$1.go -imports=".=github.com/tomcraven/gotable"
}

create_mock output
create_mock column
