class AddMike2Table < ActiveRecord::Migration[7.0]
  def change
    create_table :mike2 do |t|
      t.string :name
      t.string :email
      t.timestamps
    end
  end
end
