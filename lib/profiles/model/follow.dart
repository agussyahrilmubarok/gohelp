import 'package:shelf_dart_api/common/model/base_entity.dart';

class Follow extends BaseEntity {
  final String followerId;
  final String followeeId;

  Follow(
      {required String id,
      required this.followerId,
      required this.followeeId,
      required DateTime createdAt,
      required DateTime updatedAt,
      required DateTime? deletedAt})
      : super(
            id: id,
            createdAt: createdAt,
            updatedAt: updatedAt,
            deletedAt: deletedAt);
}
