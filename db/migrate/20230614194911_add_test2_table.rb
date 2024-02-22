class AddTest2Table < ActiveRecord::Migration[7.0]
  def change
    create_table :test2 do |t|
      t.string :name
      t.string :description
      t.timestamps
    end
  end
end
