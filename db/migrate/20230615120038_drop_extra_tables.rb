class DropExtraTables < ActiveRecord::Migration[7.0]
  def change
    drop_table :examples
    drop_table :mike
    drop_table :mike2
    drop_table :test
    drop_table :test2
    drop_table :test3
    drop_table :test4
    drop_table :test5
  end
end
