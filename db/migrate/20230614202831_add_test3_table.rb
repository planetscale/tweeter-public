class AddTest3Table < ActiveRecord::Migration[7.0]
  def change
    create_table :test3 do |t|
      t.string :name
      t.timestamps
    end
  end
end
